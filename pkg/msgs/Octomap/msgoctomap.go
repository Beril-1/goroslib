package Octomap

import (
	"github.com/aler9/goroslib/pkg/msg"
	"github.com/aler9/goroslib/pkg/msgs/std_msgs"
)

type OccupancyGrid struct {
	msg.Package `ros:"octomap_msgs"`
	Header      std_msgs.Header
  Binary      bool
  Id          string
  Resolution  float64
	Data        []int8
}
