package opium

type Token string

type Room struct {
	Name string
	Questions []Question

	QuestionLimits map[User] int
}

type User struct {
	Id uint32
	Token Token

	Name string
}

type Question struct {
	Id uint32
	Text string
	Likes uint32
	UsersLiked []*User
}

func (q *Question ) GotLikedBy(u *User){
	if !IsInList(q.UsersLiked, u){
		q.Likes += 1
		q.UsersLiked = append(q.UsersLiked, u)
	}
}

func (room *Room) AddQuestion(q Question){
	room.Questions = append(room.Questions, q)
}

func (room *Room) HandleNewUser(){

}