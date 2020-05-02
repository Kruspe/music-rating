import { useQuery } from 'react-query';
import { Auth } from 'aws-amplify';

const getUserId = async () => {
  const { id } = await Auth.currentUserInfo();
  return id;
};

const getToken = async () => {
  const currentSession = await Auth.currentSession();
  return currentSession.getAccessToken().getJwtToken();
};

const useUser = () => {
  const userId = useQuery('userId', getUserId);
  const token = useQuery('token', getToken);
  return { userId, token };
};

export default useUser;
