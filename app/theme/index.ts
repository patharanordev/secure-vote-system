import { createTheme } from '@mui/material/styles';

export const theme = createTheme({
    components: {
        MuiCssBaseline: {
            styleOverrides: {
                body: {
                    margin: 0,
                    padding: 0,
                    width: '100%'
                }
            }
        }
    }
});