package main

import (
	"machine"
	"time"
)

func main() {
	machine.Serial.Configure(machine.UARTConfig{})
	time.Sleep(10 * time.Second) // Give USB time to connect

	println("TinyGo PWM Benchmark Starting...")
	println("Benchmark: PWM Setup")
	println("task,method,setup_time_us")

	const pwmPin = machine.GPIO15 // GPIO15 = physical pin 20

	pwm := machine.PWM7 // GPIO15 maps to PWM7

	start := time.Now()

	err := pwm.Configure(machine.PWMConfig{
		Period: 8200, // ~122 kHz for audible buzzer tone
	})
	if err != nil {
		println("pwm,error,", err.Error())
		return
	}

	ch, err := pwm.Channel(pwmPin)
	if err != nil {
		println("pwm,error,", err.Error())
		return
	}

	pwm.Set(ch, pwm.Top()/2) // 50% duty cycle

	elapsed := time.Since(start).Microseconds()
	println("pwm,setup,", elapsed)

	for {
		time.Sleep(time.Second)
	}
}
