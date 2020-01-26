import React, { useState } from 'react';
import { Tabs, Tab } from '@material-ui/core';
import Overview from './overview';
import Rating from './rating';
import EstimateWacken from './estimation/wacken';

const TabBar = () => {
  const [tab, setTab] = useState(0);
  return (
    <>
      <Tabs
        value={tab}
        indicatorColor="primary"
        variant="fullWidth"
        onChange={(event, value) => setTab(value)}
      >
        <Tab label="Overview" />
        <Tab label="Rate" />
        <Tab label="Estimate Wacken" />
      </Tabs>
      {tab === 0 && <Overview />}
      {tab === 1 && <Rating />}
      {tab === 2 && <EstimateWacken />}
    </>
  );
};

export default TabBar;
