import React, { useContext } from 'react';
import UserContext from '../../context/UserContext';

const Overview = () => {
  const { ratedBands } = useContext(UserContext);
  return (
    <>
      {ratedBands && <p>OverviewPage</p>}
    </>
  );
};

export default Overview;
