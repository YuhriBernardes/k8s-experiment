import { FC, useState } from "react";
import Typography from "@material-ui/core/Typography";
import Card from "@material-ui/core/Card";
import CardHeader from "@material-ui/core/CardHeader";
import CardContent from "@material-ui/core/CardContent";
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Button from '@material-ui/core/Button';
import Toolbar from '@material-ui/core/Toolbar';
import Replay from '@material-ui/icons/Replay'

type User ={
  id: number,
  name: string,
  age: number
}

const UserList: FC = ()=>{
  const [rows, setRows] = useState<User[]>([])

  const onClick = ()=>{
    setRows([...rows, {id: 1, name: "Yuhri", age: 12}])
  }

  return (
    <Card>
      <CardHeader title={
        <Toolbar>
          <Typography style={{marginRight: '1rem'}} variant="h3">Users</Typography>
          <Button onClick={onClick} variant="contained" color="secondary"><Replay/></Button>
        </Toolbar>
      }/>
      <CardContent>
        <Table >
          <TableHead>
            <TableRow>
              <TableCell>ID</TableCell>
              <TableCell>Name</TableCell>
              <TableCell>Age</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {rows.map((u, i)=>(
            <TableRow>
              <TableCell> {i}</TableCell>
              <TableCell>{u.name}</TableCell>
              <TableCell>{u.age}</TableCell>
            </TableRow>
            ))}
          </TableBody>
        </Table>
        
      </CardContent>
    </Card>
  )
}

export default UserList
