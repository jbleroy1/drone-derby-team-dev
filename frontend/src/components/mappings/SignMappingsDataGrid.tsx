/* eslint-disable react/jsx-key */
import * as React from 'react';
import Box from '@mui/material/Box';
import EditIcon from '@mui/icons-material/Edit';
import SaveIcon from '@mui/icons-material/Save';
import CancelIcon from '@mui/icons-material/Close';
import {
  GridToolbar,
  GridRowModesModel,
  GridRowModes,
  DataGrid,
  GridColumns,
  GridRowParams,
  MuiEvent,
  GridActionsCellItem,
  GridEventListener,
  GridRowId,
  GridRowModel,
} from '@mui/x-data-grid';
import getCollection from '../../firebase/getCollection';
import updateDocument from '../../firebase/updateDocument';

export default function SignMappingsDataGrid() {
  const [rows, setRows] = React.useState<any[]>([]);
  getCollection(setRows, rows, 'SignMappings');

  const [rowModesModel, setRowModesModel] = React.useState<GridRowModesModel>({});

  const handleRowEditStart = (params: GridRowParams, event: MuiEvent<React.SyntheticEvent>) => {
    event.defaultMuiPrevented = true;
  };

  const handleRowEditStop: GridEventListener<'rowEditStop'> = (params, event) => {
    event.defaultMuiPrevented = true;
  };

  const handleEditClick = (id: GridRowId) => () => {
    setRowModesModel({ ...rowModesModel, [id]: { mode: GridRowModes.Edit } });
  };

  const handleSaveClick = (id: GridRowId) => () => {
    setRowModesModel({ ...rowModesModel, [id]: { mode: GridRowModes.View } });
  };

  const handleCancelClick = (id: GridRowId) => () => {
    setRowModesModel({
      ...rowModesModel,
      [id]: { mode: GridRowModes.View, ignoreModifications: true },
    });

    const editedRow = rows.find((row) => row.id === id);
    if (editedRow!.isNew) {
      setRows(rows.filter((row) => row.id !== id));
    }
  };

  const processRowUpdate = (newRow: GridRowModel) => {
    const updatedRow = { ...newRow, isNew: false };
    setRows(rows.map((row) => (row.id === newRow.id ? updatedRow : row)));

    type signMapping = {
      name: string;
      pitch: string;
      roll: string;
      yawDirection: string;
      yawRotation: number;
    };

    const data: signMapping = {
      name: newRow.name,
      pitch: newRow.pitch,
      roll: newRow.roll,
      yawDirection: newRow.yawDirection,
      yawRotation: newRow.yawRotation,
    };
    updateDocument('SignMappings', newRow.id, data);
    return updatedRow;
  };

  const columns: GridColumns = [
    {
      field: 'imgURL',
      headerName: 'Sign',
      width: 180,
      editable: false,
      renderCell: (params) => <img src={params.value} alt="sign mapping" width="30px" />,
    },
    {
      field: 'yawDirection',
      headerName: 'Yaw Direction',
      type: 'singleSelect',
      valueOptions: ['Left', 'Right'],
      flex: 1,
      editable: true,
    },
    {
      field: 'yawRotation',
      headerName: 'Yaw Rotation (°)',
      type: 'number',
      flex: 1,
      headerAlign: 'left',
      align: 'left',
      editable: true,
    },
    {
      field: 'roll',
      headerName: 'Roll Direction',
      type: 'singleSelect',
      valueOptions: ['Left', 'Right'],
      flex: 1,
      editable: true,
    },
    {
      field: 'pitch',
      headerName: 'Pitch Direction',
      type: 'singleSelect',
      valueOptions: ['Forward', 'Backward'],
      flex: 1,
      editable: true,
    },
    {
      field: 'actions',
      type: 'actions',
      headerName: 'Actions',
      width: 100,
      cellClassName: 'actions',
      getActions: ({ id }) => {
        const isInEditMode = rowModesModel[id]?.mode === GridRowModes.Edit;

        if (isInEditMode) {
          return [
            <GridActionsCellItem icon={<SaveIcon />} label="Save" onClick={handleSaveClick(id)} />,
            <GridActionsCellItem
              icon={<CancelIcon />}
              label="Cancel"
              className="textPrimary"
              onClick={handleCancelClick(id)}
              color="inherit"
            />,
          ];
        }

        return [
          <GridActionsCellItem
            icon={<EditIcon />}
            label="Edit"
            className="textPrimary"
            onClick={handleEditClick(id)}
            color="inherit"
          />,
        ];
      },
    },
  ];

  return (
    <Box
      sx={{
        height: 1000,
        width: '100%',
        '& .actions': {
          color: 'text.secondary',
        },
        '& .textPrimary': {
          color: 'text.primary',
        },
      }}
    >
      <DataGrid
        initialState={{
          sorting: {
            sortModel: [{ field: 'imgURL', sort: 'desc' }],
          },
        }}
        rows={rows}
        columns={columns}
        editMode="row"
        rowModesModel={rowModesModel}
        onRowModesModelChange={(newModel) => setRowModesModel(newModel)}
        onRowEditStart={handleRowEditStart}
        onRowEditStop={handleRowEditStop}
        processRowUpdate={processRowUpdate}
        components={{
          Toolbar: GridToolbar,
        }}
        componentsProps={{
          toolbar: { setRows, setRowModesModel },
        }}
        experimentalFeatures={{ newEditingApi: true }}
        rowsPerPageOptions={[20]}
        disableSelectionOnClick
      />
    </Box>
  );
}
