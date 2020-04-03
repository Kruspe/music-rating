import React, { useContext, useEffect, useState } from 'react';
import { Storage } from 'aws-amplify';
import { Grid } from '@material-ui/core';
import UserContext from '../../context/UserContext';
import EstimationCard from '../../components/estimationCard';

const EstimateWacken = () => {
  const [bandsToBeRated, setBandsToBeRated] = useState([]);
  const { ratedArtists } = useContext(UserContext);
  useEffect(() => {
    const getWackenBands = async () => {
      const wackenBandsUrl = await Storage.get('wacken.json');
      const response = await fetch(wackenBandsUrl, { headers: { 'Content-Type': 'application/json' } });
      return response.json();
    };

    const getUnratedWackenBands = async () => {
      const wackenBands = await getWackenBands();
      const ratedArtistNames = ratedArtists.map((ratedArtist) => ratedArtist.band);
      setBandsToBeRated(wackenBands
        .filter((wackenBand) => !ratedArtistNames.includes(wackenBand.artist)));
    };
    getUnratedWackenBands();
  }, [ratedArtists]);

  return (
    <Grid container>
      {bandsToBeRated.map((bandToBeRated) => (
        <Grid key={bandToBeRated.artist} item xs={12} sm={12} md={4} lg={3} xl={2}>
          <EstimationCard
            artist={bandToBeRated.artist}
            image={bandToBeRated.image}
            submitCallback={(artist) => {
              setBandsToBeRated(
                bandsToBeRated.filter((band) => band.artist !== artist),
              );
            }}
          />
        </Grid>
      ))}
    </Grid>
  );
};

export default EstimateWacken;
