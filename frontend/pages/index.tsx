import { Container, makeStyles, Theme } from '@material-ui/core'
import Head from 'next/head'
import Form from '../components/Form'
import AppBar from '@material-ui/core/AppBar'
import Toolbar from '@material-ui/core/Toolbar'
import Typography from '@material-ui/core/Typography'
import CssBaseline from '@material-ui/core/CssBaseline'
import UserList from 'components/UserList'

const useStyles = makeStyles({
  root: {
    backgroundColor: "#F8F9FE"
  }
})

export default function Home() {
  const classes = useStyles()
  return (
    <>
      <Head>
        <title>My kube</title>
      </Head>

      <CssBaseline/>
      <AppBar style={{marginBottom: '1.5rem', padding: '1rem 1.5rem'}} position="static">
        <Toolbar>
          <Typography variant="h3">My Kube</Typography>
        </Toolbar>
      </AppBar>
      <div className={classes.root} >
        <Container maxWidth="sm">
          <Form/>
          <UserList/>
        </Container>
      </div>
    </>
  )
}
