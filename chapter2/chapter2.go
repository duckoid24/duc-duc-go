package main

import (
	"context"
	"fmt"
)
type UserType = int
type SubscriptionType = int

const (
   unknownUserType UserType = iota
   lead
   customer
   churned
   lostLead
)

const (
   unknownSubscriptionType SubscriptionType = iota
   basic
   premium
   exclusive
)

type LeadRequest struct {
	email string
 }
 type Lead struct {
	id string
 }
 type LeadCreator interface {
	CreateLead(ctx context.Context, request LeadRequest) (Lead, error)
 }
 type Customer struct {
	leadID string
	userID string
 }
 func (c *Customer) UserID() string {
	return c.userID
 }
 func (c *Customer) SetUserID(userID string) {
	c.userID = userID
 }
 type LeadConvertor interface {
	Convert(ctx context.Context, subSelection SubscriptionType) (Customer, error)
 }
 func (l Lead) Convert(ctx context.Context, subSelection SubscriptionType) (Customer, error) {
	//TODO implement me
	panic("implement me")
 }

func main() {
	fmt.Println(unknownUserType)
	fmt.Println(lead)
	fmt.Println(customer)
	fmt.Println(churned)
	fmt.Println(lostLead)
}

/*
When a lead uses our app for the first time, 
they must pick one of three subscription plans. 
These are basic, premium, and exclusive. 
Depending on which they pick determines which features 
they get access to within the app. This may change over time.
Once a subscription plan has been created, we consider that 
the lead has converted to a customer, and we call them a customer 
until they churn. At this point, we call them a lead again. 
After 6 months, we call them a lost lead and we might target them 
with a re-engagement campaign, which could include a discount code. 
Once a plan is created, we set up a recurring payment to 
capture funds from the customer by direct debit.
*/