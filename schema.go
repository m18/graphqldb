package main

const schema string = `
scalar Time # non-standard but predefined by graphql-go

schema {
	query: Query
}

type Query {
	orders(first: Int!): [Order!]!,
}

type Order {
	id: Int!,
	customer: Customer!,
	time: Time!,
	products: [OrderProduct!]!,
}

type Customer {
	id: Int!,
	name: String!,
}

type OrderProduct {
	id: Int!,
	name: String!,
	price: Float!,
	quantity: Int!,
	totalPrice: Float!,
}
`
