import { useQuery } from 'react-query';
import { API, graphqlOperation } from 'aws-amplify';
import { listRatings } from '../../graphql/queries';

const fetchAllRatings = () => API.graphql(graphqlOperation(listRatings));

const useRating = () => useQuery('ratedArtists', fetchAllRatings);

export default useRating;
