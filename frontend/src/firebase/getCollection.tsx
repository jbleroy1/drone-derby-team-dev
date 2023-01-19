import { collection, DocumentData, onSnapshot, query } from 'firebase/firestore';
// eslint-disable-next-line no-unused-vars
import React, { useEffect, useState } from 'react';
import { db } from './config';

const getCollection = (setRows: any, rows: any[], collectionName: string) => {
  useEffect(() => {
    const q = query(collection(db, collectionName));
    const unsubscribe = onSnapshot(
      q,
      (snapshot) => {
        const docs: any[] = [];
        snapshot.forEach((doc: DocumentData) => {
          docs.push({
            id: doc.id,
            ...doc.data(),
          });
        });
        setRows(docs);
      },
      (error) => {
        alert(error.message);
        console.log(error);
      }
    );
    return () => unsubscribe();
  }, [collectionName]);

  return { rows };
};

export default getCollection;
