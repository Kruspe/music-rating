import { useQuery } from 'react-query';
import { API, Auth } from 'aws-amplify';

const getUserId = async () => {
  const { id } = await Auth.currentUserInfo();
  return id;
};

const getToken = async () => {
  const currentSession = await Auth.currentSession();
  return currentSession.getAccessToken().getJwtToken();
};

const getRatedArtists = async (key, userId, token) => API.get('musicrating',
  `/api/v1/ratings/bands/${userId}`,
  { header: { Authorization: `Bearer ${token}` } });

const useUser = () => {
  const { data: userId } = useQuery('userId', getUserId);
  const { data: token } = useQuery('token', getToken);
  const { data: ratedArtists = [] } = useQuery(userId && token && ['ratedArtists', userId, token], getRatedArtists);
  return { userId, token, ratedArtists };
};

export default useUser;
