package stage

import (
	"log"
	"time"

	"pl.home/game2/conf"
	el "pl.home/game2/element"
	"pl.home/game2/remote"
	"pl.home/game2/remote/encoder"
	"pl.home/game2/remote/message"
)

type MotionController struct {
	paddleLeft  *el.Paddle
	paddleRight *el.Paddle
	ball        *el.Ball

	tickerReset chan struct{}
	tickerStop  chan struct{}

	start time.Time
}

func NewMotionController(paddleLeft *el.Paddle, paddleRight *el.Paddle, ball *el.Ball) *MotionController {
	motion := &MotionController{
		paddleLeft:  paddleLeft,
		paddleRight: paddleRight,
		ball:        ball,

		tickerReset: make(chan struct{}),
		tickerStop:  make(chan struct{}),
	}
	// go motion.list()
	// go motion.send()

	return motion
}

func (mc *MotionController) Start() {
	mc.gameSpeedUp()
	mc.start = time.Now()
}

func (mc *MotionController) Stop() {
	mc.tickerStop <- struct{}{}
}

func (mc *MotionController) gameSpeedUp() {
	tick := time.Second * 10
	ticker := time.NewTicker(tick)
	go func() {
		for {
			select {
			case <-ticker.C:
				mc.ball.Speed++
				mc.paddleLeft.Speed++
				mc.paddleRight.Speed++
				log.Println("speed up!")
			case <-mc.tickerReset:
				ticker.Reset(tick)
			// log.Println("speed up reset")
			case <-mc.tickerStop:
				ticker.Stop()
				return
			}
		}
	}()

}

func (mc *MotionController) Reset() {
	mc.tickerReset <- struct{}{}
	mc.start = time.Now()
}

func (mc *MotionController) Update(keyUp, keyDown, keyW, keyS bool) {
	if keyUp && mc.paddleRight.Y > 0 {
		mc.paddleRight.Y -= mc.paddleRight.Speed
	}
	if keyDown && mc.paddleRight.Y < conf.ScreenHeight-conf.PaddleHeight {
		mc.paddleRight.Y += mc.paddleRight.Speed
	}

	if keyW && mc.paddleLeft.Y > 0 {
		mc.paddleLeft.Y -= mc.paddleLeft.Speed
	}
	if keyS && mc.paddleLeft.Y < conf.ScreenHeight-conf.PaddleHeight {
		mc.paddleLeft.Y += mc.paddleLeft.Speed
	}
	if mc.start.UnixMilli() < time.Now().Add(-time.Millisecond*300).UnixMilli() {
		if mc.ball.DirectRight {
			mc.ball.X += mc.ball.Speed
		} else {
			mc.ball.X -= mc.ball.Speed
		}
		if mc.ball.DirectBottom {
			mc.ball.Y += mc.ball.Speed
		} else {
			mc.ball.Y -= mc.ball.Speed
		}
	}
}

func (mc *MotionController) list() {
	mess := make(chan []byte)
	lst, err := remote.NewListener(conf.PortListener, mess)
	if err != nil {
		panic(err)
	}
	for {
		msg := <-mess
		encr := encoder.NewBinaryEncoder[message.PaddleMsg]()
		paddle, err := encr.Decode([]byte(msg))
		if err != nil {
			continue
			// panic(err)
		}
		// fmt.Printf("BALL: %f\n", paddle.Y)
		if conf.Player == "left" {
			mc.paddleRight.Y = paddle.Y
		} else {
			mc.paddleLeft.Y = paddle.Y
		}
	}
	lst.Close()
}

func (mc *MotionController) send() {
	sdr, err := remote.NewSender(conf.PortSender)
	if err != nil {
		panic(err)
	}
	id := 0
	for {
		id++
		y := mc.paddleRight.Y
		if conf.Player == "left" {
			y = mc.paddleLeft.Y
		}

		paddle := message.PaddleMsg{
			Y: y,
		}
		encr := encoder.NewBinaryEncoder[message.PaddleMsg]()
		enc, err := encr.Encode(paddle)
		if err != nil {
			panic(err)
		}
		if err := sdr.Send(enc); err != nil {
			// log.Println("problem podczas wyslania wiadomosci")
		}
		time.Sleep(time.Millisecond * 50)
	}
}
