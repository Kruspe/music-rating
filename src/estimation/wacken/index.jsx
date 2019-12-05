import React, { useEffect, useState } from 'react';
import { API, Auth, Storage } from 'aws-amplify';
import Rating from '../../rating';

const EstimateWacken = () => {
  const [bandsToBeRated, setBandsToBeRated] = useState([]);
  useEffect(() => {
    const getWackenBands = async () => {
      const wackenBandsUrl = await Storage.get('wacken.json');
      const response = await fetch(wackenBandsUrl, { headers: { 'Content-Type': 'application/json' } });
      return response.json();
    };
    const getRatedBands = async () => {
      const currentSession = await Auth.currentSession();
      const currentUserInfo = await Auth.currentUserInfo();
      const token = currentSession.getAccessToken().getJwtToken();
      const ratedBands = await API.get('musicrating', `/bands/${currentUserInfo.id}`, { header: { Authorization: `Bearer ${token}` } });
      return ratedBands.map((ratedBand) => ratedBand.band);
    };

    const getUnratedWackenBands = async () => {
      const wackenBands = await getWackenBands();
      const ratedBands = await getRatedBands();
      setBandsToBeRated(wackenBands.filter((wackenBand) => !ratedBands.includes(wackenBand)));
    };
    getUnratedWackenBands();
  }, []);

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
