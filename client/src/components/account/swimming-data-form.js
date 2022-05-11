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
  import {getUserActions} from "../../redux/users/actions"
  import {updateUserSwimmingData} from "../../redux/swimming-data/action"
  import { useEffect, useState } from 'react';
import { errorMessageAlertMessage, successGeneralMesssage } from 'src/utils/sweerAlert';
  
  
  
  
   const SwimmingDataForm = (props) => {
  
    const swimmingType = [
        {
            name: "Butterfly",
            values:  "Butterfly",
            id:1

        },
        {
            name: "Backstroke",
            values:  "Backstroke",
            id:2
        },
        {
            name: "Breaststroke",
            values:  "Breaststroke",
            id:3
        },
        {
            name: "Freestyle",
            values:  "Freestyle",
            id:4
        },

    ]


    const { usersList, getUserActions, updateSwimmerAction, updateUserSwimmingData } = props
  
    const getUsers = async () => {
       await getUserActions();
      
     }
     useEffect(
      () => {
        getUsers();
    }, [updateSwimmerAction])
    
  
  
    const [values, setValues] = useState({
      total_distance_covered: 0,
      stroke_count: 0,
      heart_rate: 0,
      time_taken_in_seconds: 0,
      swimming_type: "",
      user_id: ""
    });
  
    const handleChange = (event) => {
      setValues({
        ...values,
        [event.target.name]: event.target.value
      });
    };
  
    const handleSubmit = async (event) => {
        try {
            event.preventDefault();
           
            if (values.swimming_type  && values.user_id && values.total_distance_covered ){
                
                const formData = {
                    total_distance_covered: parseInt(values.total_distance_covered),
                    stroke_count: parseInt(values.stroke_count),
                    heart_rate: parseInt(values.heart_rate),
                    time_taken_in_seconds:parseInt(values.time_taken_in_seconds),
                    swimming_type: values.swimming_type,
                    user_id:values.user_id,   
                }
                
                await updateUserSwimmingData(formData)
                successGeneralMesssage();
                setValues({
                    total_distance_covered: 0,
                    stroke_count: 0,
                    heart_rate: 0,
                    time_taken_in_seconds: 0,
                    swimming_type: "",
                    user_id: ""
                })
            }else {
                errorMessageAlertMessage("User and Swimming type must be selected")
            }
          
        } catch (error) {
           
            errorMessageAlertMessage(error.response)
        }
      
    
    }
  
    return (
      <form
        autoComplete="off"
        noValidate
      >
        <Card>
          <CardHeader
            subheader="The information can be edited"
            title="Add Swimming Data"
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
                  label="Total Distance Covered"
                  name="total_distance_covered"
                  type="number"
                  onChange={handleChange}
                  required
                  value={values.total_distance_covered}
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
                  label="Stroke Count"
                  name="stroke_count"
                  type="number"
                  onChange={handleChange}
                  required
                  value={values.stroke_count}
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
                  label="Heart Beat rate"
                  name="heart_rate"
                  type="number"
                  onChange={handleChange}
                  required
                  value={values.heart_rate}
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
                  label="Time taken"
                  name="time_taken_in_seconds"
                  onChange={handleChange}
                  type="number"
                  value={values.time_taken_in_seconds}
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
                  label="User"
                  name="user_id"
                  onChange={handleChange}
                  required
                  select
                  SelectProps={{ native: true }}
                  value={values.user_id}
                  variant="outlined"
                >
                  <option
                      value={""}
                    >
                     Select
                    </option>
                  { usersList ? usersList?.map((option) => (
                    <option
                      key={option.id}
                      value={option.id}
                    >
                      {option?.firstname + " " + option?.surname}
                    </option>
                  )): null}
                </TextField>
              </Grid>
              <Grid
                item
                md={6}
                xs={12}
              >
                <TextField
                  fullWidth
                  label="SwimmingType"
                  name="swimming_type"
                  onChange={handleChange}
                  required
                  select
                  SelectProps={{ native: true }}
                  value={values.swimming_type}
                  variant="outlined"
                >
                    <option
                      value={""}
                    >
                     Select
                    </option>
                  {swimmingType?.map((option) => (
                    <option
                      key={option.id}
                      value={option.values}
                    >
                      {option.name}
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
    return state.users;
  }
  
  
   export default connect(mapStateToProps, {updateUserSwimmingData, getUserActions})(SwimmingDataForm)