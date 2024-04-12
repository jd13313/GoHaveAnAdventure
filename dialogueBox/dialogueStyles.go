package dialogueBox

type DialogueStyle struct {
	CornerTopLeft     string
	CornerTopRight    string
	Side              string
	DividerLeft       string
	DividerRight      string
	Line              string
	CornerBottomLeft  string
	CornerBottomRight string
}

var DoubleLines = DialogueStyle{
	CornerTopLeft:     "╔",
	CornerTopRight:    "╗",
	Side:              "║",
	DividerLeft:       "╠",
	DividerRight:      "╣",
	Line:              "═",
	CornerBottomLeft:  "╚",
	CornerBottomRight: "╝",
}

var BasicLines = DialogueStyle{
	CornerTopLeft:     "+",
	CornerTopRight:    "+",
	Side:              "|",
	DividerLeft:       "+",
	DividerRight:      "+",
	Line:              "-",
	CornerBottomLeft:  "+",
	CornerBottomRight: "+",
}

var ThinLines = DialogueStyle{
	CornerTopLeft:     "┌",
	CornerTopRight:    "┐",
	Side:              "│",
	DividerLeft:       "├",
	DividerRight:      "┤",
	Line:              "─",
	CornerBottomLeft:  "└",
	CornerBottomRight: "┘",
}

var CurvyLines = DialogueStyle{
	CornerTopLeft:     "╭",
	CornerTopRight:    "╮",
	Side:              "│",
	DividerLeft:       "├",
	DividerRight:      "┤",
	Line:              "─",
	CornerBottomLeft:  "╰",
	CornerBottomRight: "╯",
}
