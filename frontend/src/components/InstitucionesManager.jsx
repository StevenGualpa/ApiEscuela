import { useState, useEffect } from 'react';
import axios from 'axios';
import ConfirmDialog from './ConfirmDialog';

const InstitucionesManager = ({ onBack }) => {
  const API_BASE_URL = 'http://localhost:3000';
  
  const [instituciones, setInstituciones] = useState([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState('');
  const [success, setSuccess] = useState('');
  const [showForm, setShowForm] = useState(false);
  const [editingItem, setEditingItem] = useState(null);
  const [formData, setFormData] = useState({});
  const [showConfirmDialog, setShowConfirmDialog] = useState(false);
  const [itemToDelete, setItemToDelete] = useState(null);

  useEffect(() => {
    loadInstituciones();
  }, []);

  const loadInstituciones = async () => {
    setLoading(true);
    try {
      const response = await axios.get(`${API_BASE_URL}/instituciones`);
      setInstituciones(response.data);
    } catch (err) {
      setError('Error al cargar instituciones: ' + (err.response?.data?.error || err.message));
    } finally {
      setLoading(false);
    }
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    setLoading(true);
    setError('');
    setSuccess('');

    try {
      if (editingItem) {
        await axios.put(`${API_BASE_URL}/instituciones/${editingItem.ID}`, formData);
        setSuccess('Institución actualizada exitosamente');
      } else {
        await axios.post(`${API_BASE_URL}/instituciones`, formData);
        setSuccess('Institución creada exitosamente');
      }
      
      await loadInstituciones();
      resetForm();
    } catch (err) {
      setError('Error al guardar: ' + (err.response?.data?.error || err.message));
    } finally {
      setLoading(false);
    }
  };

  const handleEdit = (item) => {
    setEditingItem(item);
    setFormData({ ...item });
    setShowForm(true);
  };

  const handleDelete = (item) => {
    setItemToDelete(item);
    setShowConfirmDialog(true);
  };

  const handleConfirmDelete = async () => {
    if (itemToDelete) {
      try {
        await axios.delete(`${API_BASE_URL}/instituciones/${itemToDelete.ID}`);
        setSuccess('Institución eliminada exitosamente');
        await loadInstituciones();
      } catch (err) {
        setError('Error al eliminar: ' + (err.response?.data?.error || err.message));
      }
    }
    setShowConfirmDialog(false);
    setItemToDelete(null);
  };

  const handleCancelDelete = () => {
    setShowConfirmDialog(false);
    setItemToDelete(null);
  };

  const resetForm = () => {
    setFormData({});
    setEditingItem(null);
    setShowForm(false);
  };

  const handleInputChange = (e) => {
    const { name, value } = e.target;
    setFormData(prev => ({
      ...prev,
      [name]: value
    }));
  };

  return (
    <div className="min-h-screen bg-gray-100">
      {/* Header */}
      <header className="bg-white shadow">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="flex justify-between items-center py-6">
            <div className="flex items-center space-x-4">
              <button
                onClick={onBack}
                className="bg-gray-600 hover:bg-gray-700 text-white p-2 rounded-full transition-colors duration-200"
                title="Volver al Dashboard"
              >
                <svg className="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M15 19l-7-7 7-7" />
                </svg>
              </button>
              <h1 className="text-3xl font-bold text-gray-900">
                Gestión de Instituciones Educativas
              </h1>
            </div>
            <button
              onClick={() => setShowForm(!showForm)}
              className="bg-blue-600 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"
            >
              {showForm ? 'Cancelar' : 'Nueva Institución'}
            </button>
          </div>
        </div>
      </header>

      {/* Main Content */}
      <main className="max-w-7xl mx-auto py-6 sm:px-6 lg:px-8">
        <div className="px-4 py-6 sm:px-0">
          
          {/* Mensajes */}
          {error && (
            <div className="mb-4 bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded">
              {error}
            </div>
          )}
          
          {success && (
            <div className="mb-4 bg-green-100 border border-green-400 text-green-700 px-4 py-3 rounded">
              {success}
            </div>
          )}

          {/* Formulario */}
          {showForm && (
            <div className="bg-white shadow rounded-lg mb-6">
              <div className="px-6 py-4 border-b border-gray-200">
                <h3 className="text-lg font-medium text-gray-900">
                  {editingItem ? 'Editar Institución' : 'Nueva Institución'}
                </h3>
              </div>
              <form onSubmit={handleSubmit} className="p-6">
                <div className="grid grid-cols-1 gap-6">
                  <div>
                    <label htmlFor="nombre" className="block text-sm font-medium text-gray-700">
                      Nombre de la Institución
                    </label>
                    <input
                      type="text"
                      id="nombre"
                      name="nombre"
                      value={formData.nombre || ''}
                      onChange={handleInputChange}
                      required
                      className="mt-1 block w-full border-gray-300 rounded-md shadow-sm focus:ring-blue-500 focus:border-blue-500"
                    />
                  </div>
                  
                  <div>
                    <label htmlFor="autoridad" className="block text-sm font-medium text-gray-700">
                      Autoridad
                    </label>
                    <input
                      type="text"
                      id="autoridad"
                      name="autoridad"
                      value={formData.autoridad || ''}
                      onChange={handleInputChange}
                      className="mt-1 block w-full border-gray-300 rounded-md shadow-sm focus:ring-blue-500 focus:border-blue-500"
                    />
                  </div>
                  
                  <div>
                    <label htmlFor="contacto" className="block text-sm font-medium text-gray-700">
                      Contacto
                    </label>
                    <input
                      type="text"
                      id="contacto"
                      name="contacto"
                      value={formData.contacto || ''}
                      onChange={handleInputChange}
                      className="mt-1 block w-full border-gray-300 rounded-md shadow-sm focus:ring-blue-500 focus:border-blue-500"
                    />
                  </div>
                  
                  <div>
                    <label htmlFor="direccion" className="block text-sm font-medium text-gray-700">
                      Dirección
                    </label>
                    <textarea
                      id="direccion"
                      name="direccion"
                      value={formData.direccion || ''}
                      onChange={handleInputChange}
                      rows={3}
                      className="mt-1 block w-full border-gray-300 rounded-md shadow-sm focus:ring-blue-500 focus:border-blue-500"
                    />
                  </div>
                </div>
                
                <div className="flex justify-end space-x-3 mt-6">
                  <button
                    type="button"
                    onClick={resetForm}
                    className="bg-gray-300 hover:bg-gray-400 text-gray-800 font-bold py-2 px-4 rounded"
                  >
                    Cancelar
                  </button>
                  <button
                    type="submit"
                    disabled={loading}
                    className="bg-blue-600 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded disabled:opacity-50"
                  >
                    {loading ? 'Guardando...' : editingItem ? 'Actualizar' : 'Crear'}
                  </button>
                </div>
              </form>
            </div>
          )}

          {/* Tabla de Instituciones */}
          <div className="bg-white shadow rounded-lg">
            <div className="px-6 py-4 border-b border-gray-200">
              <h3 className="text-lg font-medium text-gray-900">Lista de Instituciones</h3>
            </div>
            <div className="overflow-x-auto">
              {loading ? (
                <div className="text-center py-4">
                  <p>Cargando...</p>
                </div>
              ) : (
                <table className="min-w-full divide-y divide-gray-200">
                  <thead className="bg-gray-50">
                    <tr>
                      <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                        ID
                      </th>
                      <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                        Nombre
                      </th>
                      <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                        Autoridad
                      </th>
                      <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                        Contacto
                      </th>
                      <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                        Dirección
                      </th>
                      <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                        Acciones
                      </th>
                    </tr>
                  </thead>
                  <tbody className="bg-white divide-y divide-gray-200">
                    {instituciones.map((institucion) => (
                      <tr key={institucion.ID}>
                        <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                          {institucion.ID}
                        </td>
                        <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                          {institucion.nombre}
                        </td>
                        <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                          {institucion.autoridad || 'N/A'}
                        </td>
                        <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                          {institucion.contacto || 'N/A'}
                        </td>
                        <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                          {institucion.direccion || 'N/A'}
                        </td>
                        <td className="px-6 py-4 whitespace-nowrap text-sm font-medium space-x-2">
                          <button
                            onClick={() => handleEdit(institucion)}
                            className="inline-flex items-center px-3 py-1.5 bg-blue-500 hover:bg-blue-600 text-white text-sm font-medium rounded-md transition-colors duration-200 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
                          >
                            <svg className="w-4 h-4 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z" />
                            </svg>
                            Editar
                          </button>
                          <button
                            onClick={() => handleDelete(institucion)}
                            className="inline-flex items-center px-3 py-1.5 bg-red-500 hover:bg-red-600 text-white text-sm font-medium rounded-md transition-colors duration-200 focus:outline-none focus:ring-2 focus:ring-red-500 focus:ring-offset-2"
                          >
                            <svg className="w-4 h-4 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
                            </svg>
                            Eliminar
                          </button>
                        </td>
                      </tr>
                    ))}
                  </tbody>
                </table>
              )}
            </div>
          </div>
        </div>
      </main>

      {/* Modal de confirmación para eliminación */}
      <ConfirmDialog
        isOpen={showConfirmDialog}
        onClose={handleCancelDelete}
        onConfirm={handleConfirmDelete}
        title="Confirmar Eliminación"
        message={`¿Está seguro de que desea eliminar la institución "${itemToDelete?.nombre || 'esta institución'}"? Esta acción no se puede deshacer.`}
        confirmText="Eliminar"
        cancelText="Cancelar"
      />
    </div>
  );
};

export default InstitucionesManager;
