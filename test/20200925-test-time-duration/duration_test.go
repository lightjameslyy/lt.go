package main

import (
	"testing"
	"time"
)

func Test_ParseDuration(t *testing.T) {
	for _, timeStr := range []string{
		"50.744296106s",
		"8.113665794s",
		"7.801632785s",
		"4.781600786s",
		"9.863410583s",
		"7.098224208s",
		"6.988115666s",
		"47.987236362s",
		"1m13.393615886s",
		"2m5.262120637s",
		"5m27.126051309s",
		"7m24.266986235s",
		"9m3.262048517s",
		"10m12.278089239s",
		"10m56.569478287s",
		"11m48.532706132s",
		"12m7.995335092s",
		"12m47.832810917s",
		"47m20.621328057s",
	} {
		dura, _ := time.ParseDuration(timeStr)
		t.Logf("duration: %.2f", float64(dura.Microseconds())/1000000)
	}
}
