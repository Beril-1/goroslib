//nolint:golint
package vision_msgs

import (
	"github.com/aler9/goroslib/pkg/msg"
	"github.com/aler9/goroslib/pkg/msgs/sensor_msgs"
	"github.com/aler9/goroslib/pkg/msgs/std_msgs"
)

type Detection2D struct {
	msg.Package `ros:"vision_msgs"`
	Header      std_msgs.Header
	Results     []ObjectHypothesisWithPose
	Bbox        BoundingBox2D
	SourceImg   sensor_msgs.Image
	IsTracking  bool
	TrackingId  string
}
