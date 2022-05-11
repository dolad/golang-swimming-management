import {
  Box,
  Button,
  Card,
  CardContent,
  CardHeader,
  Divider,
  Grid,
  TextField
} from '@mui/material';
import { connect } from 'react-redux';
import {getSwimmerProfileAction, updateSwimmerAction} from "../../redux/users/actions"
import { useEffect, useState } from 'react';




 const SwimmerProfileDetails = (props) => {

  const states = require("../../utils/uk-cities.json")

  const {swimmerProfile, getSwimmerProfileAction, updateSwimmerAction } = props

  const getSwimmerProfile = async () => {
    const response = await getSwimmerProfileAction();
    return response.data;
   }
   useEffect(
    () => {
      getSwimmerProfile();
  }, [updateSwimmerAction])
  


  const [values, setValues] = useState({
    firstname: swimmerProfile?.firstname,
    surname: swimmerProfile?.surname,
    address: swimmerProfile?.address,
    phone: swimmerProfile?.phonenumber,
    state: swimmerProfile?.state,
    country: swimmerProfile?.country
  });

  const handleChange = (event) => {
    setValues({
      ...values,
      [event.target.name]: event.target.value
    });
  };

  const handleSubmit = async (event) => {
     event.preventDefault();
     updateSwimmerAction(values);

  }

  return (
    <form
      autoComplete="off"
      noValidate
    >
      <Card>
        <CardHeader
          subheader="The information can be edited"
          title="Profile"
        />
        <Divider />
        <CardContent>
          <Grid
            container
            spacing={3}
          >
            <Grid
              item
              md={6}
              xs={12}
            >
              <TextField
                fullWidth
                helperText="Please specify the first name"
                label="First name"
                name="firstname"
                onChange={handleChange}
                required
                value={values.firstname}
                variant="outlined"
              />
            </Grid>
            <Grid
              item
              md={6}
              xs={12}
            >
              <TextField
                fullWidth
                label="Last name"
                name="surname"
                onChange={handleChange}
                required
                value={values.surname}
                variant="outlined"
              />
            </Grid>
            <Grid
              item
              md={6}
              xs={12}
            >
              <TextField
                fullWidth
                label="Address"
                name="address"
                onChange={handleChange}
                required
                value={values.address}
                variant="outlined"
              />
            </Grid>
            <Grid
              item
              md={6}
              xs={12}
            >
              <TextField
                fullWidth
                label="Phone Number"
                name="phone"
                onChange={handleChange}
                type="number"
                value={values.phone}
                variant="outlined"
              />
            </Grid>
            <Grid
              item
              md={6}
              xs={12}
            >
              <TextField
                fullWidth
                label="Country"
                name="country"
                onChange={handleChange}
                required
                value={values.country}
                variant="outlined"
              />
            </Grid>
            <Grid
              item
              md={6}
              xs={12}
            >
              <TextField
                fullWidth
                label="Select State"
                name="state"
                onChange={handleChange}
                required
                select
                SelectProps={{ native: true }}
                value={values.state}
                variant="outlined"
              >
                {states?.map((option) => (
                  <option
                    key={option.id}
                    value={option.en_name}
                  >
                    {option.en_name}
                  </option>
                ))}
              </TextField>
            </Grid>
          </Grid>
        </CardContent>
        <Divider />
        <Box
          sx={{
            display: 'flex',
            justifyContent: 'flex-end',
            p: 2
          }}
        >
          <Button
            color="primary"
            variant="contained"
            onClick={handleSubmit}
          >
            Save details
          </Button>
        </Box>
      </Card>
    </form>
  );
};

const mapStateToProps = (state) => {
  return state.user;
}


 export default connect(mapStateToProps, {getSwimmerProfileAction, updateSwimmerAction})(SwimmerProfileDetails)