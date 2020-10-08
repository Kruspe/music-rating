import { useQuery } from 'react-query';
import { API, Auth } from 'aws-amplify';
import useUser from '../useUser';

const getRatedArtists = async (key, userId) => {
  const currentSession = await Auth.currentSession();
  const token = currentSession.getAccessToken().getJwtToken();
  return API.get('musicrating',
    `/api/v1/ratings/bands/${userId}`,
    { header: { Authorization: `Bearer ${token}` } });
};

const useRating = () => {
  const { userId } = useUser();
  return useQuery('ratedArtists', (key) => getRatedArtists(key, userId.data), { enabled: userId.data });
};

export default useRating;
