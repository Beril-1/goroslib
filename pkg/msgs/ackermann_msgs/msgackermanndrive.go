//nolint:golint
package ackermann_msgs

import (
	"github.com/aler9/goroslib/pkg/msg"
)

type AckermannDrive struct {
	msg.Package           `ros:"ackermann_msgs"`
	SteeringAngle         float32
	SteeringAngleVelocity float32
	Speed                 float32
	Acceleration          float32
	Jerk                  float32
}
