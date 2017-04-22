package main

import (
	"fmt"
	"math"

	"github.com/go-gl/mathgl/mgl32"
)

type Camera struct {
	Position          mgl32.Vec3
	Front             mgl32.Vec3
	Up                mgl32.Vec3
	FOV               float64
	speed             float32
	delta             float32
	sensitivity       float64
	isFirstMouseEvent bool
	lastX             float64
	lastY             float64
	pitch             float64
	yaw               float64
}

func NewCamera(posVec, frontVec, upVec mgl32.Vec3) *Camera {
	return &Camera{
		Position:          posVec,
		Front:             frontVec,
		Up:                upVec,
		FOV:               45.0,
		speed:             0.01,
		sensitivity:       0.25,
		isFirstMouseEvent: true,
	}
}

func NewDefaultCamera() *Camera {
	return &Camera{
		Position:          mgl32.Vec3{0.0, 0.0, 3.0},
		Front:             mgl32.Vec3{0.0, 0.0, -1.0},
		Up:                mgl32.Vec3{0.0, 1.0, 0.0},
		FOV:               45.0,
		speed:             0.01,
		sensitivity:       0.25,
		isFirstMouseEvent: true,
	}
}

func (c *Camera) SetSpeed(s float32) {
	c.speed = s
}
func (c *Camera) SetDelta(d float32) {
	c.delta = d
}
func (c *Camera) SetSensitivity(s float64) {
	c.sensitivity = s
}
func (c *Camera) MoveForward() {
	c.Position = c.Position.Add(c.Front.Mul(c.speed * c.delta))
}

func (c *Camera) MoveBackward() {
	c.Position = c.Position.Sub(c.Front.Mul(c.speed * c.delta))
}

func (c *Camera) MoveLeft() {
	c.Position = c.Position.Sub(c.Front.Cross(c.Up).Normalize().Mul(c.speed * c.delta))
}

func (c *Camera) MoveRight() {
	c.Position = c.Position.Add(c.Front.Cross(c.Up).Normalize().Mul(c.speed * c.delta))
}

func (c *Camera) CurrentView() mgl32.Mat4 {
	return mgl32.LookAtV(c.Position, c.Position.Add(c.Front), c.Up)
}

func (c *Camera) HandleCursorEvent(xpos, ypos float64) {
	if c.isFirstMouseEvent {
		c.lastX = xpos
		c.lastY = ypos
		c.isFirstMouseEvent = false
	}

	xoffset := xpos - c.lastX
	yoffset := c.lastY - ypos
	c.lastX = xpos
	c.lastY = ypos

	xoffset = xoffset * c.sensitivity
	yoffset = yoffset * c.sensitivity

	c.yaw = c.yaw + xoffset
	c.pitch = c.pitch + yoffset

	if c.pitch > 89.0 {
		c.pitch = 89.0
	}

	if c.pitch < -89.0 {
		c.pitch = -89.0
	}

	fX := float32(math.Cos(float64(mgl32.DegToRad(float32(c.yaw)))) * math.Cos(float64(mgl32.DegToRad(float32(c.pitch)))))
	fY := float32(math.Sin(float64(mgl32.DegToRad(float32(c.pitch)))))
	fZ := float32(math.Sin(float64(mgl32.DegToRad(float32(c.yaw)))) * math.Cos(float64(mgl32.DegToRad(float32(c.pitch)))))

	front := mgl32.Vec3{fX, fY, fZ}
	c.Front = front.Normalize()
}

func (c *Camera) HandleScrollEvent(xoffset, yoffset float64) {
	if c.FOV >= 1.0 && c.FOV <= 45.0 {
		c.FOV = c.FOV - yoffset
	}
	if c.FOV <= 1.0 {
		c.FOV = 1.0
	}
	if c.FOV >= 45.0 {
		c.FOV = 45.0
	}
	fmt.Println(c.FOV)
}
