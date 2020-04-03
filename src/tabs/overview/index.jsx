import React, { useContext } from 'react';
import UserContext from '../../context/UserContext';

const Overview = () => {
  const { ratedArtists } = useContext(UserContext);
  return (
    <>
      {ratedArtists && <p>OverviewPage</p>}
    </>
  );
};

export default Overview;
