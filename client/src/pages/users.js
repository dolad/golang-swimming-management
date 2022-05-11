import Head from 'next/head';
import { Box, Container } from '@mui/material';
import UserListResult from '../components/customer/users-list-results';
import { CustomerListToolbar } from '../components/customer/customer-list-toolbar';
import { DashboardLayout } from '../components/dashboard-layout';
import { connect } from 'react-redux'
import {getUserActions} from '../redux/users/actions'
import { useEffect, useState } from 'react';

const Users = (props) => {
  const {getUserActions} = props;

  const getUserList = async () => {
   const response = await getUserActions();
   return response.data;
  }

  useEffect(
    () => {
      getUserList();
  }, [])
  
  return(
  <>
    <Head>
      <title>
        Users
      </title>
    </Head>
    <Box
      component="main"
      sx={{
        flexGrow: 1,
        py: 8
      }}
    >
      <Container maxWidth={false}>
        <CustomerListToolbar />
        <Box sx={{ mt: 3 }}>
            <UserListResult />
        </Box>
      </Container>
    </Box>
  </>
)
};

Users.getLayout = (page) => (
  <DashboardLayout>
    {page}
  </DashboardLayout>
);

export default connect( null, {getUserActions})(Users);

