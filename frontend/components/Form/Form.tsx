import { FC, FormEvent, useState } from "react";
import TextField from '@material-ui/core/TextField'
import FormGroup from '@material-ui/core/FormGroup'
import Button from '@material-ui/core/Button'
import Typography from "@material-ui/core/Typography";
import Card from "@material-ui/core/Card";
import CardHeader from "@material-ui/core/CardHeader";
import CardContent from "@material-ui/core/CardContent";

const Form: FC = ()=>{

  const [form, setForm] = useState({
    name: "",
    age: 0
  })

  return (
    <Card style={{marginBottom: '1rem'}}>
      <CardHeader title={<Typography variant="h3">New User</Typography>}/>
      <CardContent>
      <form onSubmit={(e: FormEvent<HTMLFormElement>)=>{
        e.preventDefault()
        console.log("deu bão")
      }}>
        <FormGroup style={{marginBottom: '1rem'}} row>
          <TextField style={{marginRight: '1rem'}} variant="outlined" label="Name" value={form.name} onChange={(e)=>setForm({...form, name: e.target.value})}/>

          <TextField  variant="outlined" label="age" value={form.age} onChange={(e)=>setForm({...form, age: Number(e.target.value)})} type="number" inputProps={{
            min: 0,
            max: 120
          }}/>
        </FormGroup>

        <Button type="submit" color="primary" variant="contained">Add</Button>
      </form>
      </CardContent>
    </Card>
  )
}

export default Form
