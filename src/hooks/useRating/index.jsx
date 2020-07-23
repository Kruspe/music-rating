import { useQuery } from 'react-query';
import { API } from 'aws-amplify';
import useUser from '../useUser';

const getRatedArtists = async (key, userId, token) => (API.get('musicrating',
  `/api/v1/ratings/bands/${userId}`,
  { header: { Authorization: `Bearer ${token}` } }));

const useRating = () => {
  const { userId, token } = useUser();
  return useQuery('ratedArtists', (key) => getRatedArtists(key, userId.data, token.data), { enabled: userId.data && token.data });
};

export default useRating;
