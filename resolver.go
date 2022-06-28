package contactFriend

import (
	"context"
	"github.com/huyhuy1020/contactFriend/models"
)

var contacts = []*models.Contact{
	{
		ID:          "1",
		Name:        "1st contact",
		Description: "a description",
		UserID:      "1",
	},
	{
		ID:          "2",
		Name:        "2nd contact",
		Description: "second description",
		UserID:      "1",
	},
}

var users = []*models.User{
	{
		ID:       "1",
		Username: "Huy",
		Email:    "huy@gmail.com",
	},
	{
		ID:       "2",
		Username: "tuan",
		Email:    "tuan@gmai.com",
	},
}

type Resolver struct{}

func (r *Resolver) Contact() ContactfriendResolver {
	//TODO implement me
	panic("implement me")
}

func (r *Resolver) User() UserResolver {
	//TODO implement me
	panic("implement me")
}

/*func (r *Resolver) Contact() UserResolver {
	return &userResolver{r}
}*/

/*func (r *Resolver) User() UserResolver {
	//TODO implement me
	panic("implement me")
}*/

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Contacts(ctx context.Context) ([]*models.Contact, error) {
	panic("implement me")
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}

type contactResolver struct{ *Resolver }

type mutationResolver struct{ *Resolver }

func (m *mutationResolver) CreateContactFriend(ctx context.Context, input NewContactFriend) (*models.Contact, error) {
	panic("implement me")
}

func (r *contactResolver) User(ctx context.Context) ([]*models.Contact, error) {
	panic("implement me")
}

type userResolver struct{ *Resolver }

func (r *Resolver) ContactFriends() *contactResolver {
	return &contactResolver{r}
}

func (r *Resolver) User1() *userResolver {
	return &userResolver{r}
}

func (m *mutationResolver) CreateContactsFriend(ctx context.Context) ([]*models.Contact, error) {
	panic("implement me")
}

func (r *queryResolver) ContactFriend(ctx context.Context) ([]*models.Contact, error) {
	return contacts, nil
}
