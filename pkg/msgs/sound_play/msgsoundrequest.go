//nolint:golint
package sound_play

import (
	"github.com/aler9/goroslib/pkg/msg"
)

const (
	SoundRequest_BACKINGUP              int8 = 1
	SoundRequest_NEEDS_UNPLUGGING       int8 = 2
	SoundRequest_NEEDS_PLUGGING         int8 = 3
	SoundRequest_NEEDS_UNPLUGGING_BADLY int8 = 4
	SoundRequest_NEEDS_PLUGGING_BADLY   int8 = 5
	SoundRequest_ALL                    int8 = -1
	SoundRequest_PLAY_FILE              int8 = -2
	SoundRequest_SAY                    int8 = -3
	SoundRequest_PLAY_STOP              int8 = 0
	SoundRequest_PLAY_ONCE              int8 = 1
	SoundRequest_PLAY_START             int8 = 2
)

type SoundRequest struct {
	msg.Package     `ros:"sound_play"`
	msg.Definitions `ros:"int8 BACKINGUP=1,int8 NEEDS_UNPLUGGING=2,int8 NEEDS_PLUGGING=3,int8 NEEDS_UNPLUGGING_BADLY=4,int8 NEEDS_PLUGGING_BADLY=5,int8 ALL=-1,int8 PLAY_FILE=-2,int8 SAY=-3,int8 PLAY_STOP=0,int8 PLAY_ONCE=1,int8 PLAY_START=2"`
	Sound           int8
	Command         int8
	Volume          float32
	Arg             string
	Arg2            string
}
