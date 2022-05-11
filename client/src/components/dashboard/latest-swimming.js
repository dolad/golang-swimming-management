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



const LatestSwimmingData = (props) => {
  const {allUserSwimmingData} = props;
  return(
  <Card {...props}>
    <CardHeader title="Latest Swimming Datas" />
    <PerfectScrollbar>
      <Box sx={{ minWidth: 800 }}>
        <Table>
          <TableHead>
            <TableRow>
              <TableCell>
               Username
              </TableCell>
              <TableCell>
               Squad No
              </TableCell>
              <TableCell>
              TotalDistanceCovered
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
                    StrokeCount
                  </TableSortLabel>
                </Tooltip>
              </TableCell>
              <TableCell>
              HeartRate
              </TableCell>
              <TableCell>
              SwimmingType
              </TableCell>
              <TableCell>
              TimeTaken (s)
              </TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            { allUserSwimmingData && allUserSwimmingData?.map((swimmingData, index) => (
              <TableRow
                hover
                key={index +1}
              >
                <TableCell>
                  {swimmingData.User?.username}
                </TableCell>
                <TableCell>
                  {swimmingData.User?.SquadID}
                </TableCell>
                <TableCell>
                  {swimmingData.TotalDistanceCovered}
                </TableCell>
                <TableCell>
                  {swimmingData.StrokeCount}
                </TableCell>
                <TableCell>
                {swimmingData.HeartRate}
                </TableCell>
                <TableCell>
                  <SeverityPill
                    color={(swimmingData.SwimmingType === 'Freestyle' && 'success')
                    || (swimmingData.SwimmingType === 'Backstroke' && 'error')
                    || 'warning'}
                  >
                    {swimmingData.SwimmingType}
                  </SeverityPill>
                </TableCell>
                <TableCell>
                {swimmingData.TimeTakenInSeconds}
                </TableCell>
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
  return state.swimmingdata
}

export default connect( mapStateToProps, {})(LatestSwimmingData);
