import Head from 'next/head';
import { Box, Container, Grid } from '@mui/material';
import { Swimmmer } from '../components/dashboard/swimmer';
import  LatestSwimmingData  from '../components/dashboard/latest-swimming';
import { Sales } from '../components/dashboard/sales';
import { Parents } from '../components/dashboard/tasks-progress';
import { Coaches } from '../components/dashboard/coaches';
import { TotalUsers } from '../components/dashboard/total-profit';
import { TrafficByDevice } from '../components/dashboard/traffic-by-device';
import { DashboardLayout } from '../components/dashboard-layout';
import {getUserActions} from '../redux/users/actions';
import { connect } from 'react-redux';
import { useEffect, useState } from 'react';
import { getAllUsersSwimmingData } from '../redux/swimming-data/action';
import {getSquadsData} from '../redux/squad/action'
import { isAdmin } from 'src/utils/authHelper';
import SquadDetails from  "../components/dashboard/squad-detail"

const Dashboard = (props) => {
  const {user, getUserActions, getAllUsersSwimmingData, getSquadsData } = props;
  const { usersList: users, swimmers, coaches, parents } = user;

  const init = async () => {
     await getUserActions();
     await getAllUsersSwimmingData();
     await getSquadsData();
   }

   console.log(isAdmin);
 
   useEffect(
     () => {
      init();
   }, [])
   

  return(
  <>
    <Head>
      <title>
        Swimming CMS
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
        <Grid
          container
          spacing={3}
        >
          {
            isAdmin ?  <Grid
            item
            lg={3}
            sm={6}
            xl={3}
            xs={12}
          >
            <Swimmmer value={swimmers?.length} />
          </Grid> : null
          }
         {
           isAdmin ?  <Grid
           item
           xl={3}
           lg={3}
           sm={6}
           xs={12}
         >
           <Coaches value={coaches?.length}/>
         </Grid> : null
         }
          {
             isAdmin ?   <Grid
             item
             xl={3}
             lg={3}
             sm={6}
             xs={12}
           >
             <Parents value={parents?.length} />
           </Grid> : null
          }
         {
            isAdmin ? <Grid
            item
            xl={3}
            lg={3}
            sm={6}
            xs={12}
          >
            <TotalUsers value={users?.length} sx={{ height: '100%' }} />
          </Grid> : null
         }
          {
            isAdmin ? null :  <Grid
            item
            lg={8}
            md={12}
            xl={9}
            xs={12}
          >
            <Sales />
          </Grid> 
          }
         
         {
            isAdmin ? null :  <Grid
            item
            lg={4}
            md={6}
            xl={3}
            xs={12}
          >
            <TrafficByDevice sx={{ height: '100%' }} />
          </Grid>
          }
           
          <Grid
            item
            lg={12}
            md={12}
            xl={12}
            xs={12}
          >
            <LatestSwimmingData />
          </Grid>
          <Grid
            item
            lg={12}
            md={12}
            xl={12}
            xs={12}
          >
             <SquadDetails /> 
          </Grid>
        </Grid>
      </Container>
    </Box>
  </>
)};

Dashboard.getLayout = (page) => (
  <DashboardLayout>
    {page}
  </DashboardLayout>
);


const mapStateToProps = (state) => {
  console.log(state.users);
  return {
    user:state.users,
    auth:state.auth
    }
}

export default connect( mapStateToProps, {getUserActions, getAllUsersSwimmingData, getSquadsData})(Dashboard);




export async function  getStaticProps(context) {
  return {
    props: {
      protected: true,
      userType: "admin"
    }
  }
}