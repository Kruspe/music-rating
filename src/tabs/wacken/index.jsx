import React, { useEffect, useState } from 'react';
import { Storage } from 'aws-amplify';
import { Grid } from '@material-ui/core';
import { useQuery } from 'react-query';
import EstimationCard from '../../components/estimationCard';
import useRating from '../../hooks/useRating';

const getWackenBands = async () => {
  const wackenBandsUrl = await Storage.get('wacken.json');
  const response = await fetch(wackenBandsUrl, { headers: { 'Content-Type': 'application/json' } });
  return response.json();
};

const Wacken = () => {
  const [bandsToBeRated, setBandsToBeRated] = useState([]);
  const { data: ratedArtists } = useRating();
  const { data: wackenBands } = useQuery('wackenBands', getWackenBands);

  useEffect(() => {
    if (ratedArtists && wackenBands) {
      const ratedArtistNames = ratedArtists.map((ratedArtist) => ratedArtist.band);
      setBandsToBeRated(wackenBands
        .filter((wackenBand) => !ratedArtistNames.includes(wackenBand.artist)));
    }
  }, [ratedArtists, wackenBands]);

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

export default Wacken;
