import { useQuery } from 'react-query';
import { API } from 'aws-amplify';
import useUser from '../useUser';

const getRatedArtists = async (key, userId, token) => (API.get('musicrating',
  `/api/v1/ratings/bands/${userId}`,
  { header: { Authorization: `Bearer ${token}` } }));

const useRating = () => {
  const { userId, token } = useUser();
  return useQuery(userId.data && token.data && 'ratedArtists', [userId.data, token.data], getRatedArtists);
};

export default useRating;
