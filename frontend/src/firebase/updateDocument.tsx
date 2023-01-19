import { updateDoc, doc } from 'firebase/firestore';
import { db } from './config';

const updateDocument = async (collection: string, documentId: string, data: any) => {
  const docRef = doc(db, collection, documentId);
  await updateDoc(docRef, data).catch((error) => {
    console.error(error);
  });
};

export default updateDocument;
