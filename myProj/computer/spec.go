package computer

type Spec struct { //exported struct  
    Maker string //exported field
    Price int //exported field
    model string //unexported field

}