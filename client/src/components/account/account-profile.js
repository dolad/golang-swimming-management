import {
  Avatar,
  Box,
  Button,
  Card,
  CardActions,
  CardContent,
  Divider,
  Typography
} from '@mui/material';

import {getSwimmerProfileAction} from "../../redux/users/actions"
import { connect } from 'react-redux';

const user = {
  avatar: '/static/images/avatars/avatar_6.png',
  city: 'Los Angeles',
  country: 'USA',
  jobTitle: 'Senior Developer',
  name: 'Katarina Smith',
  timezone: 'GTM-7'
};

const AccountProfile = (props) => {
  
  const {swimmerProfile} = props
  

  return (
  <Card >
    <CardContent>
      <Box
        sx={{
          alignItems: 'center',
          display: 'flex',
          flexDirection: 'column'
        }}
      >
        <Avatar
          src={'/static/images/avatars/avatar_6.png'}
          sx={{
            height: 64,
            mb: 2,
            width: 64
          }}
        />
        <Typography
          color="textPrimary"
          gutterBottom
          variant="h5"
        >
          { `${swimmerProfile?.surname} ${swimmerProfile?.firstname}` } 
        </Typography>
        <Typography
          color="textSecondary"
          variant="body2"
        >
          {`${swimmerProfile?.state} ${swimmerProfile?.country}`}
        </Typography>
        <Typography
          color="textSecondary"
          variant="body2"
        >
          {swimmerProfile?.role?.name}
        </Typography>
      </Box>
    </CardContent>
    <Divider />
  </Card>
)
};

const mapStateToProps = (state) => {
  return state.users;
}

export default connect(mapStateToProps, {getSwimmerProfileAction})(AccountProfile)