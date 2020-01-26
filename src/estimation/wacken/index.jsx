import React, { useEffect, useState, useContext } from 'react';
import { Storage } from 'aws-amplify';
import Rating from '../../rating';
import UserContext from '../../context/UserContext';

const EstimateWacken = () => {
  const [bandsToBeRated, setBandsToBeRated] = useState([]);
  const { ratedBands } = useContext(UserContext);
  useEffect(() => {
    const getWackenBands = async () => {
      const wackenBandsUrl = await Storage.get('wacken.json');
      const response = await fetch(wackenBandsUrl, { headers: { 'Content-Type': 'application/json' } });
      return response.json();
    };

    const getUnratedWackenBands = async () => {
      const wackenBands = await getWackenBands();
      const ratedBandNames = ratedBands.map((ratedBand) => ratedBand.band);
      setBandsToBeRated(wackenBands.filter((wackenBand) => !ratedBandNames.includes(wackenBand)));
    };
    getUnratedWackenBands();
  }, [ratedBands]);

  return (
    <>
      {bandsToBeRated.map((bandToBeRated) => (
        <Rating
          key={bandToBeRated}
          bandName={bandToBeRated}
          onSubmitBehaviour={() => setBandsToBeRated(
            bandsToBeRated.filter((band) => band !== bandToBeRated),
          )}
        />
      ))}
    </>
  );
};

export default EstimateWacken;
