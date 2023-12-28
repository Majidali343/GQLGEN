package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.41

import (
	"context"
	"gqlgen-jokes-2/graph/model"
)

// Jokes is the resolver for the Jokes field.
func (r *queryResolver) Jokes(ctx context.Context) ([]*model.Joke, error) {

	JokesData := []*model.Joke{
		{Text: "Why don't scientists trust atoms? Because they make up everything!", ID: "5"},
		{Text: "What did one plate say to another plate? Tonight, dinner's on me!", ID: "8"},
		{Text: "Why did the scarecrow win an award?Because he was outstanding in his field!", ID: "6"},
		{Text: "Parallel lines have so much in common.It's a shame they'll never meet.", ID: "6"},
		{Text: "What do you call a fish wearing a crown?A kingfish.", ID: "6"},
		{Text: "Why did the bicycle fall over?Because it was two-tired!", ID: "6"},
		{Text: "What do you call fake spaghetti? An impasta.", ID: "6"},
		{Text: "I only know 25 letters of the alphabet. I don't know y.", ID: "6"},
		{Text: "How do you organize a space party?You planet!", ID: "6"},
	}

	return JokesData, nil
	// panic(fmt.Errorf("not implemented: Jokes - Jokes"))
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
