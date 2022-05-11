import { useState } from 'react';
import PerfectScrollbar from 'react-perfect-scrollbar';
import PropTypes from 'prop-types';
import { format } from 'date-fns';
import { connect } from 'react-redux';
import {getUserActions} from '../../redux/users/actions'


import {
  Avatar,
  Box,
  Card,
  Checkbox,
  Table,
  TableBody,
  TableCell,
  TableHead,
  TablePagination,
  TableRow,
  Typography
} from '@mui/material';
import { getInitials } from '../../utils/get-initials';

const UserListResult = (props) => {

  const {usersList : users} = props;
  const [selectedUserId, setselectedUserId] = useState([]);
  const [limit, setLimit] = useState(10);
  const [page, setPage] = useState(0);

  const handleSelectAll = (event) => {
    let newselectedUserId;

    if (event.target.checked) {
      newselectedUserId = users?.map((user) => user.id);
    } else {
      newselectedUserId = [];
    }

    setselectedUserId(newselectedUserId);
  };

  const handleSelectOne = (event, id) => {
    const selectedIndex = selectedUserId.indexOf(id);
    let newselectedUserId = [];

    if (selectedIndex === -1) {
      newselectedUserId = newselectedUserId.concat(selectedUserId, id);
    } else if (selectedIndex === 0) {
      newselectedUserId = newselectedUserId.concat(selectedUserId.slice(1));
    } else if (selectedIndex === selectedUserId.length - 1) {
      newselectedUserId = newselectedUserId.concat(selectedUserId.slice(0, -1));
    } else if (selectedIndex > 0) {
      newselectedUserId = newselectedUserId.concat(
        selectedUserId.slice(0, selectedIndex),
        selectedUserId.slice(selectedIndex + 1)
      );
    }

    setselectedUserId(newselectedUserId);
  };

  const handleLimitChange = (event) => {
    setLimit(event.target.value);
  };

  const handlePageChange = (event, newPage) => {
    setPage(newPage);
  };

  return (
    <Card >
      <PerfectScrollbar>
        <Box sx={{ minWidth: 1050 }}>
          <Table>
            <TableHead>
              <TableRow>
                <TableCell padding="checkbox">
                  <Checkbox
                    checked={selectedUserId.length === users?.length}
                    color="primary"
                    indeterminate={
                      selectedUserId.length > 0
                      && selectedUserId.length < users?.length
                    }
                    onChange={handleSelectAll}
                  />
                </TableCell>
                <TableCell>
                  Name
                </TableCell>
                <TableCell>
                  Email
                </TableCell>
                <TableCell>
                  Location
                </TableCell>
                <TableCell>
                  Phone
                </TableCell>
                <TableCell>
                  Date of birth
                </TableCell>
                <TableCell>
                  Role
                </TableCell>
                <TableCell>
                  Registration date
                </TableCell>
              </TableRow>
            </TableHead>
            <TableBody>

              {users && users.length > 0 && users?.slice(0, limit).map((user) => (
                <TableRow
                  hover
                  key={user.id}
                  selected={selectedUserId.indexOf(user.id) !== -1}
                >
                  <TableCell padding="checkbox">
                    <Checkbox
                      checked={selectedUserId.indexOf(user.id) !== -1}
                      onChange={(event) => handleSelectOne(event, user.id)}
                      value="true"
                    />
                  </TableCell>
                  <TableCell>
                    <Box
                      sx={{
                        alignItems: 'center',
                        display: 'flex'
                      }}
                    >
                      <Avatar
                        src={'/static/images/avatars/avatar_3.png'}
                        sx={{ mr: 2 }}
                      >
                        {getInitials(user.surname)}
                      </Avatar>
                      <Typography
                        color="textPrimary"
                        variant="body1"
                      >
                        {user.username}
                      </Typography>
                    </Box>
                  </TableCell>
                  <TableCell>
                    {user.email}
                  </TableCell>
                  <TableCell>
                    {`${user.postcode}, ${user.address}`}
                  </TableCell>
                  <TableCell>
                    {user.phonenumber}
                  </TableCell>
                  <TableCell>
                    {format( Date.parse(user.dateofbirth), 'dd/MM/yyyy')}
                  </TableCell>
                  <TableCell>
                    {user.Role.name}
                  </TableCell>
                  <TableCell>
                    {format( Date.parse(user.created_at), 'dd/MM/yyyy')}
                  </TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </Box>
      </PerfectScrollbar>
      <TablePagination
        component="div"
        count={users?.length}
        onPageChange={handlePageChange}
        onRowsPerPageChange={handleLimitChange}
        page={page}
        rowsPerPage={limit}
        rowsPerPageOptions={[5, 10, 25]}
      />
    </Card>
  );
};


const mapStateToProps = (state) => {
  console.log(state.users);
  return state.users
}
export default connect(mapStateToProps, null)(UserListResult);

