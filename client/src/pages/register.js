import Head from 'next/head';
import NextLink from 'next/link';
import { useRouter } from 'next/router';
import { useFormik } from 'formik';
import * as Yup from 'yup';
import moment from 'moment';
import {
  Box,
  Button,
  Checkbox,
  Container, FormControl,
  FormHelperText, InputLabel,
  Link, MenuItem, Select,
  TextField,
  Typography
} from '@mui/material';
import ArrowBackIcon from '@mui/icons-material/ArrowBack';
import { connect } from 'react-redux';
import {registerAction} from "../redux/authenticated/action"
import { DatePicker, LocalizationProvider } from '@mui/lab';
import AdapterDateFns from '@mui/lab/AdapterDateFns';
import { makeStyles } from '@mui/styles';

const useStyles = makeStyles(() => ({
  textField: {
    padding: 5,
    marginBottom: 10,
  },
  passwordInput:{
    marginBottom: 10,
  },
  dateInput:{
    marginBottom: 10,
  },
  input: {
    color: "white"
  }
}));

const Register = (props) => {
  const router = useRouter();
  const classes = useStyles();

  const formik = useFormik({
    initialValues: {
      email: '',
      username: '',
      firstname: '',
      surname: '',
      phonenumber: '',
      address: '',
      postcode: '',
      role: '',
      dateofbirth:'',
      password: ''
    },
    validationSchema: Yup.object({
      email: Yup
        .string()
        .email(
          'Must be a valid email')
        .max(250)
        .required(
          'Email is required'),
      username: Yup
        .string()
        .max(250)
        .required(
          'username  is required'),
      firstname: Yup
        .string()
        .max(250)
        .required(
          'First name is required'),
      surname: Yup
        .string()
        .max(250)
        .required(
          'Surname is required'),
      phonenumber: Yup.string()
                      .max(255)
                      .required("Phone number is required")
                      .matches('^(?:0|\\+?44)(?:\\d\\s?){9,10}$', "Invalid phone number"),
      postcode: Yup.string()
                      .max(255)
                      .required("Invalid postcode"),
      address: Yup.string()
        .max(255)
        .required('Address  is required'),
      dateofbirth:Yup.date().required(),

      role: Yup
        .string().optional(),
      password: Yup
        .string()
        .max(255)
        .required(
          'Password is required')

    }),
    onSubmit: async (values) => {
      try {
         await props.registerAction(values);
         router.push('/login');
      } catch (error) {
        console.log(error.response.data); 
      }
     
      
    }
  });

  return (
    <>
      <Head>
        <title>
          Register
        </title>
      </Head>
      <Box
        component="main"
        sx={{
          alignItems: 'center',
          display: 'flex',
          flexGrow: 1,
          minHeight: '100%'
        }}
      >
        <Container maxWidth="sm">
          <Box
            sx={{
              pb: 1,
              pt: 3
            }}
          >
            <Typography
              align="center"
              color="textSecondary"
              variant="h2"
            >
              Swimming CMS
            </Typography>
          </Box>
          <form onSubmit={formik.handleSubmit}>
            <Box sx={{ my: 3 }}>
              <Typography
                color="textPrimary"
                variant="h4"
              >
                Create a new account
              </Typography>
              <Typography
                color="textSecondary"
                gutterBottom
                variant="body2"
              >
                Use your email to create a new account
              </Typography>
            </Box>
            <TextField
              error={Boolean(formik.touched.firstName && formik.errors.firstName)}
              fullWidth
              helperText={formik.touched.firstName && formik.errors.firstName}
              label="First Name"
              margin="normal"
              name="firstname"
              onBlur={formik.handleBlur}
              onChange={formik.handleChange}
              value={formik.values.firstName}
              variant="outlined"
            />

            <TextField
              error={Boolean(formik.touched.surname && formik.errors.surname)}
              fullWidth
              helperText={formik.touched.surname && formik.errors.surname}
              label=" Surname"
              margin="normal"
              name="surname"
              onBlur={formik.handleBlur}
              onChange={formik.handleChange}
              value={formik.values.surname}
              variant="outlined"
            />
            <TextField
              error={Boolean(formik.touched.username && formik.errors.username)}
              fullWidth
              helperText={formik.touched.username && formik.errors.username}
              label="Username"
              margin="normal"
              name="username"
              onBlur={formik.handleBlur}
              onChange={formik.handleChange}
              value={formik.values.username}
              variant="outlined"
            />
            <TextField
              error={Boolean(formik.touched.email && formik.errors.email)}
              fullWidth
              helperText={formik.touched.email && formik.errors.email}
              label="Email Address"
              margin="normal"
              name="email"
              onBlur={formik.handleBlur}
              onChange={formik.handleChange}
              type="email"
              value={formik.values.email}
              variant="outlined"
            />
            <TextField
              inputProps={{ inputMode: 'numeric', pattern: '[0-9]*' }}
              error={Boolean(formik.touched.phonenumber && formik.errors.phonenumber)}
              fullWidth
              helperText={formik.touched.phonenumber && formik.errors.phonenumber}
              label="Phone number"
              margin="normal"
              name="phonenumber"
              onBlur={formik.handleBlur}
              onChange={formik.handleChange}
              value={formik.values.phonenumber}
              variant="outlined"
            />
            <TextField
              error={Boolean(formik.touched.address && formik.errors.address)}
              fullWidth
              helperText={formik.touched.address && formik.errors.address}
              label="Address"
              margin="normal"
              name="address"
              onBlur={formik.handleBlur}
              onChange={formik.handleChange}
              value={formik.values.address}
              variant="outlined"
            />
            <TextField
              error={Boolean(formik.touched.postcode && formik.errors.postcode)}
              fullWidth
              helperText={formik.touched.postcode && formik.errors.postcode}
              label="postcode"
              margin="normal"
              name="postcode"
              onBlur={formik.handleBlur}
              onChange={formik.handleChange}
              value={formik.values.postcode}
              variant="outlined"
            />



            <TextField
              error={Boolean(formik.touched.password && formik.errors.password)}
              fullWidth
              helperText={formik.touched.password && formik.errors.password}
              label="Password"
              margin="normal"
              name="password"
              onBlur={formik.handleBlur}
              onChange={formik.handleChange}
              className={classes.passwordInput}
              type="password"
              value={formik.values.password}
              variant="outlined"
            />
            <FormControl fullWidth>
              <InputLabel id="demo-simple-select-label">UserType</InputLabel>
              <Select
                labelId="demo-simple-select-label"
                id="demo-simple-select"
                value={formik.values.role}
                label="UserType"
                name="role"
                onChange={formik.handleChange}
                onBlur={formik.handleBlur}
              >
                <MenuItem key={1} value={"Swimmer"}>Swimmer</MenuItem>
                <MenuItem  key={2} value={"Parent"}>Parent</MenuItem>
                <MenuItem  key={3} value={"Coach"}>Coach</MenuItem>
                <MenuItem key={4} value={"Officials"}>Officials</MenuItem>
              </Select>

            </FormControl>

            <LocalizationProvider dateAdapter={AdapterDateFns}>

              <DatePicker
                renderInput={(props) => <TextField {...props} />}
                className={classes.dateInput}
                label="Data of birth"
                value={formik.values.dateofbirth}
                onChange={ value => formik.setFieldValue('dateofbirth', value)}
              />

            </LocalizationProvider>
            <Box
              sx={{
                alignItems: 'center',
                display: 'flex',
                ml: -1
              }}
            >
            </Box>

            <Box sx={{ py: 2 }}>
              <Button
                color="primary"
                disabled={formik.isSubmitting}
                fullWidth
                size="large"
                type="submit"
                variant="contained"
              >
                Sign Up Now
              </Button>
            </Box>
            <Typography
              color="textSecondary"
              variant="body2"
            >
              Have an account?
              {' '}
              <NextLink
                href="/login"
                passHref
              >
                <Link
                  variant="subtitle2"
                  underline="hover"
                >
                  Sign In
                </Link>
              </NextLink>
            </Typography>
          </form>
        </Container>
      </Box>
    </>
  );
};


export default connect(null, {registerAction})(Register);

