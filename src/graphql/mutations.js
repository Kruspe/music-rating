/* eslint-disable */
// this is an auto generated file. This will be overwritten

export const createRating = /* GraphQL */ `
  mutation CreateRating(
    $input: CreateRatingInput!
    $condition: ModelRatingConditionInput
  ) {
    createRating(input: $input, condition: $condition) {
      id
      artist
      year
      festival
      comment
      rating
      createdAt
      updatedAt
      owner
    }
  }
`;
export const updateRating = /* GraphQL */ `
  mutation UpdateRating(
    $input: UpdateRatingInput!
    $condition: ModelRatingConditionInput
  ) {
    updateRating(input: $input, condition: $condition) {
      id
      artist
      year
      festival
      comment
      rating
      createdAt
      updatedAt
      owner
    }
  }
`;
export const deleteRating = /* GraphQL */ `
  mutation DeleteRating(
    $input: DeleteRatingInput!
    $condition: ModelRatingConditionInput
  ) {
    deleteRating(input: $input, condition: $condition) {
      id
      artist
      year
      festival
      comment
      rating
      createdAt
      updatedAt
      owner
    }
  }
`;
