import { useQuery } from 'react-query';
import { Auth } from 'aws-amplify';

const getUserId = async () => {
  const { id } = await Auth.currentUserInfo();
  return id;
};

const useUser = () => {
  const userId = useQuery('userId', getUserId);
  return { userId };
};

export default useUser;
