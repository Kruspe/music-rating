import React, { forwardRef } from 'react';
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
class Overview extends React.Component {
  constructor(props) {
    super(props);
    this.state = { ratings: undefined };
  }

  async componentDidMount() {
    const currentSession = await Auth.currentSession();
    const currentUserInfo = await Auth.currentUserInfo();
    const token = currentSession.getAccessToken().getJwtToken();

    const ratings = await API.get('musicrating', `/bands/${currentUserInfo.id}`, { header: { Authorization: `Bearer ${token}` } });
    this.setState({ ratings });
  }

  render() {
    const { ratings } = this.state;
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
  }
}

export default Overview;
