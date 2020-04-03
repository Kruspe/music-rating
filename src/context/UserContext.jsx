import React from 'react';

const UserContext = React.createContext({
  userId: '', jwtToken: '', ratedArtists: [],
});

export default UserContext;
