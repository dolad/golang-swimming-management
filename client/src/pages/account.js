import Head from 'next/head';
import { Box, Container, Grid, Typography } from '@mui/material';
import  AccountProfile from '../components/account/account-profile';
import  SwimmerProfileDetails  from '../components/account/swimmer-account-details';

import  SwimmingDataForm  from '../components/account/swimming-data-form';

import { DashboardLayout } from '../components/dashboard-layout';
import { connect } from 'react-redux'
import {getUserActions} from '../redux/users/actions'
import { useEffect, useState } from 'react';

const Account = (props) => {

  const {getUserActions} = props;

  const getUserList = async () => {
   const response = await getUserActions();
   return response.data;
  }

  useEffect(
    () => {
      getUserList();
  }, [])

 return (
  <>
    <Head>
      <title>
        Account | Material Kit
      </title>
    </Head>
    <Box
      component="main"
      sx={{
        flexGrow: 1,
        py: 8
      }}
    >
      <Container maxWidth="lg">
        <Typography
          sx={{ mb: 3 }}
          variant="h4"
        >
          Account
        </Typography>
        <Grid
          container
          spacing={3}
        >
          <Grid
            item
            lg={4}
            md={6}
            xs={12}
          >
            <AccountProfile />
          </Grid>
          <Grid
            item
            lg={8}
            md={6}
            xs={12}
            padding={1}
           
           
          >
            
            <SwimmerProfileDetails />

            <Grid
            item
            lg={12}
            md={12}
            xs={12}
            paddingTop={2}
           
          >
            <SwimmingDataForm />
            </Grid>
          </Grid>

         

        </Grid>
      </Container>
    </Box>
  </>
);
    }

Account.getLayout = (page) => (
  <DashboardLayout>
    {page}
  </DashboardLayout>
);

export default connect( null, {getUserActions})(Account);

