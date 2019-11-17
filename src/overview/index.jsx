import React, { forwardRef, useEffect, useState } from 'react';
import { API, Auth } from 'aws-amplify';
import { Rating } from '@material-ui/lab';
import MaterialTable from 'material-table';

import {
  ArrowUpward, ChevronLeft, ChevronRight, Clear, FirstPage, LastPage, Search,
} from '@material-ui/icons';

const tableIcons = {
  /* eslint-disable react/jsx-props-no-spreading */
  FirstPage: forwardRef((props, ref) => <FirstPage {...props} ref={ref} />),
  LastPage: forwardRef((props, ref) => <LastPage {...props} ref={ref} />),
  NextPage: forwardRef((props, ref) => <ChevronRight {...props} ref={ref} />),
  PreviousPage: forwardRef((props, ref) => <ChevronLeft {...props} ref={ref} />),
  ResetSearch: forwardRef((props, ref) => <Clear {...props} ref={ref} />),
  Search: forwardRef((props, ref) => <Search {...props} ref={ref} />),
  SortArrow: forwardRef((props, ref) => <ArrowUpward {...props} ref={ref} />),
  /* eslint-enable react/jsx-props-no-spreading */
};
const Overview = () => {
  const [ratings, setRatings] = useState();

  useEffect(() => {
    const getRatings = async () => {
      const currentSession = await Auth.currentSession();
      const currentUserInfo = await Auth.currentUserInfo();
      const token = currentSession.getAccessToken().getJwtToken();
      setRatings(await API.get('musicrating', `/bands/${currentUserInfo.id}`, { header: { Authorization: `Bearer ${token}` } }));
    };
    getRatings();
  }, []);

  return (
    <>
      {ratings && (
        <MaterialTable
          columns={[
            { title: 'Band', field: 'band' },
            { title: 'Festival', field: 'festival' },
            { title: 'Year', field: 'year' },
            {
              title: 'Rating',
              field: 'rating',
              render: (data) => (<Rating name={data.band} value={data.rating} readOnly />),
            }]}
          data={ratings}
          icons={tableIcons}
          options={{ showTitle: false, draggable: false, pageSize: 10 }}
        />
      )}
    </>
  );
};

export default Overview;
