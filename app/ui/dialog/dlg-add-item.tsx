"use client"

import { useEffect, useState } from 'react';
import Button from '@mui/material/Button';
import TextField from '@mui/material/TextField';
import Dialog from '@mui/material/Dialog';
import DialogActions from '@mui/material/DialogActions';
import DialogContent from '@mui/material/DialogContent';
import DialogContentText from '@mui/material/DialogContentText';
import DialogTitle from '@mui/material/DialogTitle';
import { VoteInfo } from '#/types';
import { Stack } from '@mui/material';

type DialogProps = {
    isOpen: boolean
    onSave?: Function
    onCancel?: Function
}

const DEFAULT_INFO: VoteInfo = {
    itemName: '',
    itemDescription: ''
}

export default function AddItemDialog(props: DialogProps) {
    const [open, setOpen] = useState<boolean>(props.isOpen);
    const [data, setData] = useState<VoteInfo>(DEFAULT_INFO);

    const handleClickCancel = () => {
        setOpen(false)
        if (typeof props.onCancel === 'function') {
            props.onCancel()
        }
    };

    const handleClickSave = () => {
        setOpen(false)
        if (typeof props.onSave === 'function') {
            props.onSave(data)
        }
    };

    const onChange = (e) => {
        if (data) {
            switch (e.target.id) {
                case 'itemName':
                    data.itemName = e.target.value
                    break;
                case 'itemDescription':
                    data.itemDescription = e.target.value
                    break;
                default:
                    break;
            }
    
            console.log('data change:', data)
            setData(data)
        }
    }

    useEffect(() => {
      console.log('props.isOpen:', props.isOpen)
      setOpen(props.isOpen)
    }, [props.isOpen])

    return (
        <Dialog 
            open={open} 
            onClose={handleClickCancel}
            fullWidth={true}
            maxWidth={'sm'}
        >
            <DialogTitle>Add item</DialogTitle>
            <DialogContent>
                <Stack spacing={2}>
                    <TextField
                        autoFocus
                        margin="dense"
                        id="itemName"
                        label="Item name"
                        fullWidth
                        variant="standard"
                        onChange={onChange}
                    />
                    <TextField
                        margin="dense"
                        id="itemDescription"
                        label="Item description"
                        fullWidth
                        variant="standard"
                        onChange={onChange}
                    />
                </Stack>
            </DialogContent>
            <DialogActions>
                <Button onClick={handleClickCancel}>Cancel</Button>
                <Button onClick={handleClickSave}>Save</Button>
            </DialogActions>
        </Dialog>
    );
}