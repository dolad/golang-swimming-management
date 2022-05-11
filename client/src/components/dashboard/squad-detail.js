import { format } from 'date-fns';
import { v4 as uuid } from 'uuid';
import PerfectScrollbar from 'react-perfect-scrollbar';
import {
  Box,
  Button,
  Card,
  CardHeader,
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableRow,
  TableSortLabel,
  Tooltip
} from '@mui/material';
import ArrowRightIcon from '@mui/icons-material/ArrowRight';
import { SeverityPill } from '../severity-pill';
import { connect } from 'react-redux';
import EditIcon from '@mui/icons-material/Edit';
import PreviewIcon from '@mui/icons-material/Preview';



const SquadDetails = (props) => {
  const {allSquadData} = props;
//   const [showSquadDetails, setSquadDetails] = useState(false);

//   const handleViewToggle = (data) => () => {
//     setSquadDetails(true);
//    console.log(data);
//   };


//   const ShowSquadDetails = () => (
//     <Modal
//     open={showSquadDetails}
//     onClose={handleClose}
//     aria-labelledby="modal-modal-title"
//     aria-describedby="modal-modal-description"
//   >
//     <Box sx={style}>
//       <Typography id="modal-modal-title" variant="h6" component="h2">
//         Text in a modal
//       </Typography>
//       <Typography id="modal-modal-description" sx={{ mt: 2 }}>
//         Duis mollis, est non commodo luctus, nisi erat porttitor ligula.
//       </Typography>
//     </Box>
//    </Modal>
//   )

  return(
  <Card {...props}>
    <CardHeader title="Squad Datas" />
    <PerfectScrollbar>
      <Box sx={{ minWidth: 800 }}>
        <Table>
          <TableHead>
            <TableRow>
              <TableCell>
               Squad Number
              </TableCell>
              <TableCell>
               Squad Name
              </TableCell>
              <TableCell>
              Coach Username
              </TableCell>
              <TableCell sortDirection="desc">
                <Tooltip
                  enterDelay={300}
                  title="Sort"
                >
                  <TableSortLabel
                    active
                    direction="desc"
                  >
                    Number of Swimmers
                  </TableSortLabel>
                </Tooltip>
              </TableCell>
              <TableCell>
                Actions
              </TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            { allSquadData && allSquadData?.map((squadData, index) => (
              <TableRow
                hover
                key={index +1}
              >
                <TableCell>
                  {squadData.ID}
                </TableCell>
                <TableCell>
                {squadData.name}
                </TableCell>
                <TableCell>
                 {squadData?.Coach?.username }
                </TableCell>
                <TableCell>
                {squadData?.Swimmers?.length }
                </TableCell>
                <Table>
                 {/* <PreviewIcon onClick={handleViewToggle(squadData.ID)} sx={{ color: 'green' }} /> */}
                 {/* <EditIcon onClick={handleViewToggle(squadData.ID)} sx={{ color: 'red' }} /> */}
                </Table>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </Box>
    </PerfectScrollbar>
    <Box
      sx={{
        display: 'flex',
        justifyContent: 'flex-end',
        p: 2
      }}
    >
      <Button
        color="primary"
        endIcon={<ArrowRightIcon fontSize="small" />}
        size="small"
        variant="text"
      >
        View all
      </Button>
    </Box>
  </Card>
)};

const mapStateToProps = (state) => {
  return state.squad
}

export default connect( mapStateToProps, {})(SquadDetails);
