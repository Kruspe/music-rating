import React from 'react';

const UserContext = React.createContext({
  userId: '', jwtToken: '', ratedBands: [],
});

export default UserContext;
