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

var StyleDouble = DialogueStyle{
	CornerTopLeft:     "╔",
	CornerTopRight:    "╗",
	Side:              "║",
	DividerLeft:       "╠",
	DividerRight:      "╣",
	Line:              "═",
	CornerBottomLeft:  "╚",
	CornerBottomRight: "╝",
}

var StyleBasic = DialogueStyle{
	CornerTopLeft:     "+",
	CornerTopRight:    "+",
	Side:              "|",
	DividerLeft:       "+",
	DividerRight:      "+",
	Line:              "-",
	CornerBottomLeft:  "+",
	CornerBottomRight: "+",
}

var StyleThin = DialogueStyle{
	CornerTopLeft:     "┌",
	CornerTopRight:    "┐",
	Side:              "│",
	DividerLeft:       "├",
	DividerRight:      "┤",
	Line:              "─",
	CornerBottomLeft:  "└",
	CornerBottomRight: "┘",
}

var StyleCurved = DialogueStyle{
	CornerTopLeft:     "╭",
	CornerTopRight:    "╮",
	Side:              "│",
	DividerLeft:       "├",
	DividerRight:      "┤",
	Line:              "─",
	CornerBottomLeft:  "╰",
	CornerBottomRight: "╯",
}
