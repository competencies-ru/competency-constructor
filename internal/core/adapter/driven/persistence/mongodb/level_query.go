package mongodb

import "go.mongodb.org/mongo-driver/bson"

func makeFilterUgsn(levelID string) bson.A {
	return bson.A{
		bson.D{{Key: "$match", Value: bson.D{{Key: "_id", Value: levelID}}}},
		bson.D{
			{
				Key: "$project",
				Value: bson.D{
					{
						Key: "ugsn",
						Value: bson.D{
							{
								Key: "$sortArray",
								Value: bson.D{
									{Key: "input", Value: "$ugsn"},
									{Key: "sortBy", Value: bson.D{{Key: "code", Value: 1}}},
								},
							},
						},
					},
				},
			},
		},
		bson.D{
			{
				Key: "$unset",
				Value: bson.A{
					"ugsn.specialties",
				},
			},
		},
	}
}

func makeFilterSpecialties(ugsnID string) bson.A {
	return bson.A{
		bson.D{{Key: "$match", Value: makeMatchUgsnID(ugsnID)}},
		bson.D{
			{
				Key: "$project",
				Value: bson.D{
					{
						Key: "ugsn",
						Value: bson.D{
							{
								Key: "$filter",
								Value: bson.D{
									{Key: "input", Value: "$ugsn"},
									{Key: "as", Value: "item"},
									{
										Key: "cond",
										Value: bson.D{
											{
												Key: "$eq",
												Value: bson.A{
													"$$item.id",
													ugsnID,
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		bson.D{
			{
				Key: "$project",
				Value: bson.D{
					{Key: "_id", Value: 0},
					{
						Key: "specialties",
						Value: bson.D{
							{
								Key: "$sortArray",
								Value: bson.D{
									{Key: "input", Value: "$ugsn.specialties"},
									{Key: "sortBy", Value: bson.D{{Key: "code", Value: 1}}},
								},
							},
						},
					},
				},
			},
		},
		bson.D{{Key: "$unwind", Value: "$specialties"}},
		bson.D{
			{
				Key: "$unset",
				Value: bson.A{
					"specialties.programs",
				},
			},
		},
	}
}

func makeFilterPrograms(specialtiesID string) bson.A {
	return bson.A{
		bson.D{{Key: "$match", Value: bson.D{{Key: "ugsn.specialties.id", Value: bson.D{{Key: "$eq", Value: specialtiesID}}}}}},
		bson.D{{Key: "$unwind", Value: "$ugsn"}},
		bson.D{
			{
				Key: "$project",
				Value: bson.D{
					{
						Key: "specialties",
						Value: bson.D{
							{
								Key: "$filter",
								Value: bson.D{
									{Key: "input", Value: "$ugsn.specialties"},
									{Key: "as", Value: "item"},
									{
										Key: "cond",
										Value: bson.D{
											{
												Key: "$eq",
												Value: bson.A{
													"$$item.id",
													specialtiesID,
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		bson.D{{Key: "$unwind", Value: "$specialties"}},
		bson.D{
			{
				Key: "$project",
				Value: bson.D{
					{Key: "_id", Value: 0},
					{
						Key: "programs",
						Value: bson.D{
							{
								Key: "$sortArray",
								Value: bson.D{
									{Key: "input", Value: "$specialties.programs"},
									{Key: "sortBy", Value: bson.D{{Key: "code", Value: 1}}},
								},
							},
						},
					},
				},
			},
		},
	}
}

func makeMatchUgsnID(ugsnID string) bson.D {
	return bson.D{{
		Key: "ugsn",
		Value: bson.D{{
			Key: "$elemMatch",
			Value: bson.D{{
				Key:   "id",
				Value: bson.D{{Key: "$eq", Value: ugsnID}},
			}},
		}},
	}}
}
