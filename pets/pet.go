package pets

type Pet interface {
	Feed(food string, name string) string
	GiveAttention(activity string) string
	IsHungry() bool
}
