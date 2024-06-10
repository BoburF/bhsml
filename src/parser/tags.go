package parser

type Position struct{
    StartIndex int
    EndIndex int
}

type Tag struct{
    Line int
    Position Position
    Type string
    Name string
    Atributes map[string]string
    Chidlren []Tag
}
