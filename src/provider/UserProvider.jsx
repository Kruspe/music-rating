import { API, Auth } from 'aws-amplify';
import React, { useEffect, useState } from 'react';
import PropTypes from 'prop-types';
import UserContext from '../context/UserContext';

const UserProvider = ({ children }) => {
  const [jwtToken, setJwtToken] = useState('');
  const [userId, setUserId] = useState('');
  const [ratedBands, setRatedBands] = useState([]);

  useEffect(() => {
    const getJwtToken = async () => {
      const currentSession = await Auth.currentSession();
      setJwtToken(currentSession.getAccessToken().getJwtToken());
    };
    const getUserId = async () => {
      const currentUserInfo = await Auth.currentUserInfo();
      setUserId(currentUserInfo.id);
    };
    getJwtToken();
    getUserId();
  }, []);

  useEffect(() => {
    const getRatedBands = async () => {
      setRatedBands(await API.get('musicrating',
        `/bands/${userId}`, {
          header: { Authorization: `Bearer ${jwtToken}` },
        }));
    };
    if (userId && jwtToken) {
      getRatedBands();
    }
  }, [jwtToken, userId]);

  return (
    <UserContext.Provider value={{ jwtToken, userId, ratedBands }}>
      {children}
    </UserContext.Provider>
  );
};

UserProvider.propTypes = {
  children: PropTypes.node,
};

UserProvider.defaultProps = {
  children: [],
};

export default UserProvider;
