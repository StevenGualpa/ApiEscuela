import { useState, useEffect } from 'react';
import axios from 'axios';
import ConfirmDialog from './ConfirmDialog';

const SystemConfig = ({ onBack }) => {
  // URL base de la API
  const API_BASE_URL = 'http://localhost:3000';
  
  const [currentSection, setCurrentSection] = useState('overview');
  const [provincias, setProvincias] = useState([]);
  const [ciudades, setCiudades] = useState([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState('');
  const [success, setSuccess] = useState('');
  const [showForm, setShowForm] = useState(false);
  const [editingItem, setEditingItem] = useState(null);
  const [formData, setFormData] = useState({});
  const [showConfirmDialog, setShowConfirmDialog] = useState(false);
  const [itemToDelete, setItemToDelete] = useState(null);

  // Cargar datos según la sección actual
  useEffect(() => {
    if (currentSection !== 'overview') {
      loadData();
    }
  }, [currentSection]);

  const loadData = async () => {
    setLoading(true);
    try {
      switch (currentSection) {
        case 'provincias':
          const provinciasRes = await axios.get(`${API_BASE_URL}/provincias`);
          setProvincias(provinciasRes.data);
          break;
        case 'ciudades':
          const [ciudadesRes, provinciasForCities] = await Promise.all([
            axios.get(`${API_BASE_URL}/ciudades`),
            axios.get(`${API_BASE_URL}/provincias`)
          ]);
          setCiudades(ciudadesRes.data);
          setProvincias(provinciasForCities.data);
          break;
      }
    } catch (err) {
      setError('Error al cargar los datos: ' + (err.response?.data?.error || err.message));
    } finally {
      setLoading(false);
    }
  };

  const handleInputChange = (e) => {
    setFormData({
      ...formData,
      [e.target.name]: e.target.value
    });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    setLoading(true);
    setError('');
    setSuccess('');

    try {
      const endpoint = currentSection;
      if (editingItem) {
        await axios.put(`${API_BASE_URL}/${endpoint}/${editingItem.ID}`, formData);
        setSuccess(`${getSectionTitle()} actualizado exitosamente`);
      } else {
        await axios.post(`${API_BASE_URL}/${endpoint}`, formData);
        setSuccess(`${getSectionTitle()} creado exitosamente`);
      }

      await loadData();
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
        await axios.delete(`${API_BASE_URL}/${currentSection}/${itemToDelete.ID}`);
        setSuccess('Elemento eliminado exitosamente');
        await loadData();
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

  const getSectionTitle = () => {
    switch (currentSection) {
      case 'provincias': return 'Provincia';
      case 'ciudades': return 'Ciudad';
      default: return 'Elemento';
    }
  };

  const getProvinciaName = (provinciaId) => {
    const provincia = provincias.find(p => p.ID === provinciaId);
    return provincia ? provincia.nombre : 'N/A';
  };

  const renderForm = () => {
    switch (currentSection) {
      case 'provincias':
        return (
          <div className="space-y-4">
            <div>
              <label className="block text-sm font-medium text-gray-700">Nombre de la Provincia</label>
              <input
                type="text"
                name="nombre"
                value={formData.nombre || ''}
                onChange={handleInputChange}
                required
                className="mt-1 block w-full border border-gray-300 rounded-md shadow-sm py-2 px-3 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500"
                placeholder="Ej: Cotopaxi"
              />
            </div>
          </div>
        );

      case 'ciudades':
        return (
          <div className="space-y-4">
            <div>
              <label className="block text-sm font-medium text-gray-700">Nombre de la Ciudad</label>
              <input
                type="text"
                name="ciudad"
                value={formData.ciudad || ''}
                onChange={handleInputChange}
                required
                className="mt-1 block w-full border border-gray-300 rounded-md shadow-sm py-2 px-3 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500"
                placeholder="Ej: Latacunga"
              />
            </div>
            <div>
              <label className="block text-sm font-medium text-gray-700">Provincia</label>
              <select
                name="provincia_id"
                value={formData.provincia_id || ''}
                onChange={handleInputChange}
                required
                className="mt-1 block w-full border border-gray-300 rounded-md shadow-sm py-2 px-3 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500"
              >
                <option value="">Seleccione una provincia</option>
                {provincias.map((provincia) => (
                  <option key={provincia.ID} value={provincia.ID}>
                    {provincia.nombre}
                  </option>
                ))}
              </select>
            </div>
          </div>
        );

      default:
        return null;
    }
  };

  const renderTable = () => {
    switch (currentSection) {
      case 'provincias':
        return (
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
                  Acciones
                </th>
              </tr>
            </thead>
            <tbody className="bg-white divide-y divide-gray-200">
              {provincias.map((provincia) => (
                <tr key={provincia.ID}>
                  <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                    {provincia.ID}
                  </td>
                  <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                    {provincia.nombre}
                  </td>
                  <td className="px-6 py-4 whitespace-nowrap text-sm font-medium space-x-2">
                    <button
                      onClick={() => handleEdit(provincia)}
                      className="inline-flex items-center px-3 py-1.5 bg-blue-500 hover:bg-blue-600 text-white text-sm font-medium rounded-md transition-colors duration-200 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
                    >
                      <svg className="w-4 h-4 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z" />
                      </svg>
                      Editar
                    </button>
                    <button
                      onClick={() => handleDelete(provincia)}
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
        );

      case 'ciudades':
        return (
          <table className="min-w-full divide-y divide-gray-200">
            <thead className="bg-gray-50">
              <tr>
                <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                  ID
                </th>
                <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                  Ciudad
                </th>
                <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                  Provincia
                </th>
                <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                  Acciones
                </th>
              </tr>
            </thead>
            <tbody className="bg-white divide-y divide-gray-200">
              {ciudades.map((ciudad) => (
                <tr key={ciudad.ID}>
                  <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                    {ciudad.ID}
                  </td>
                  <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                    {ciudad.ciudad}
                  </td>
                  <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                    {getProvinciaName(ciudad.provincia_id)}
                  </td>
                  <td className="px-6 py-4 whitespace-nowrap text-sm font-medium space-x-2">
                    <button
                      onClick={() => handleEdit(ciudad)}
                      className="inline-flex items-center px-3 py-1.5 bg-blue-500 hover:bg-blue-600 text-white text-sm font-medium rounded-md transition-colors duration-200 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
                    >
                      <svg className="w-4 h-4 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z" />
                      </svg>
                      Editar
                    </button>
                    <button
                      onClick={() => handleDelete(ciudad)}
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
        );

      default:
        return null;
    }
  };

  const getCurrentData = () => {
    switch (currentSection) {
      case 'provincias': return provincias;
      case 'ciudades': return ciudades;
      default: return [];
    }
  };

  return (
    <div className="min-h-screen bg-gray-100">
      {/* Header */}
      <header className="bg-white shadow">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="flex justify-between items-center py-6">
            <div className="flex items-center">
              <button 
                onClick={onBack}
                className="mr-4 text-gray-600 hover:text-gray-900"
              >
                ← Volver al Dashboard
              </button>
              <h1 className="text-3xl font-bold text-gray-900">
                Configuración del Sistema
              </h1>
            </div>
          </div>
        </div>
      </header>

      {/* Mensajes */}
      {error && (
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 pt-4">
          <div className="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded relative">
            {error}
          </div>
        </div>
      )}

      {success && (
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 pt-4">
          <div className="bg-green-100 border border-green-400 text-green-700 px-4 py-3 rounded relative">
            {success}
          </div>
        </div>
      )}

      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 pt-6">
        <div className="flex">
          {/* Sidebar */}
          <div className="w-64 bg-white shadow rounded-lg p-6 mr-6">
            <h2 className="text-lg font-semibold mb-4">Configuraciones</h2>
            <nav className="space-y-2">
              <button
                onClick={() => {
                  setCurrentSection('overview');
                  setShowForm(false);
                  resetForm();
                }}
                className={`w-full text-left px-3 py-2 rounded-md ${
                  currentSection === 'overview' ? 'bg-indigo-100 text-indigo-700' : 'text-gray-700 hover:bg-gray-100'
                }`}
              >
                Resumen
              </button>
              <button
                onClick={() => {
                  setCurrentSection('provincias');
                  setShowForm(false);
                  resetForm();
                }}
                className={`w-full text-left px-3 py-2 rounded-md ${
                  currentSection === 'provincias' ? 'bg-indigo-100 text-indigo-700' : 'text-gray-700 hover:bg-gray-100'
                }`}
              >
                Provincias
              </button>
              <button
                onClick={() => {
                  setCurrentSection('ciudades');
                  setShowForm(false);
                  resetForm();
                }}
                className={`w-full text-left px-3 py-2 rounded-md ${
                  currentSection === 'ciudades' ? 'bg-indigo-100 text-indigo-700' : 'text-gray-700 hover:bg-gray-100'
                }`}
              >
                Ciudades
              </button>
            </nav>
          </div>

          {/* Main Content */}
          <div className="flex-1">
            {currentSection === 'overview' ? (
              <div className="bg-white shadow rounded-lg p-6">
                <h2 className="text-xl font-bold mb-4">Resumen del Sistema</h2>
                <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
                  <div className="bg-blue-50 p-4 rounded-lg">
                    <div className="flex items-center">
                      <svg className="w-8 h-8 text-blue-500 mr-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M9 20l-5.447-2.724A1 1 0 013 16.382V5.618a1 1 0 011.447-.894L9 7m0 13l6-3m-6 3V7m6 10l4.553 2.276A1 1 0 0021 18.382V7.618a1 1 0 00-.553-.894L15 4m0 13V4m0 0L9 7" />
                      </svg>
                      <div>
                        <h3 className="font-semibold">Provincias</h3>
                        <p className="text-2xl font-bold text-blue-600">{provincias.length}</p>
                      </div>
                    </div>
                  </div>
                  <div className="bg-green-50 p-4 rounded-lg">
                    <div className="flex items-center">
                      <svg className="w-8 h-8 text-green-500 mr-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z" />
                        <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M15 11a3 3 0 11-6 0 3 3 0 016 0z" />
                      </svg>
                      <div>
                        <h3 className="font-semibold">Ciudades</h3>
                        <p className="text-2xl font-bold text-green-600">{ciudades.length}</p>
                      </div>
                    </div>
                  </div>
                </div>
                <div className="mt-6">
                  <p className="text-gray-600">
                    Desde aquí puedes gestionar las configuraciones básicas del sistema como provincias y ciudades. 
                    Selecciona una opción del menú lateral para comenzar.
                  </p>
                </div>
              </div>
            ) : (
              <div className="space-y-6">
                {/* Formulario */}
                {showForm && (
                  <div className="bg-white shadow rounded-lg p-6">
                    <div className="flex justify-between items-center mb-4">
                      <h2 className="text-xl font-bold">
                        {editingItem ? `Editar ${getSectionTitle()}` : `Nueva ${getSectionTitle()}`}
                      </h2>
                      <button
                        onClick={resetForm}
                        className="text-gray-500 hover:text-gray-700"
                      >
                        ✕
                      </button>
                    </div>
                    <form onSubmit={handleSubmit}>
                      {renderForm()}
                      <div className="mt-6 flex space-x-3">
                        <button
                          type="submit"
                          disabled={loading}
                          className="bg-indigo-600 hover:bg-indigo-700 text-white font-bold py-2 px-4 rounded disabled:opacity-50"
                        >
                          {loading ? 'Guardando...' : (editingItem ? 'Actualizar' : 'Crear')}
                        </button>
                        <button
                          type="button"
                          onClick={resetForm}
                          className="bg-gray-300 hover:bg-gray-400 text-gray-700 font-bold py-2 px-4 rounded"
                        >
                          Cancelar
                        </button>
                      </div>
                    </form>
                  </div>
                )}

                {/* Lista */}
                <div className="bg-white shadow rounded-lg">
                  <div className="px-6 py-4 border-b border-gray-200 flex justify-between items-center">
                    <h2 className="text-xl font-bold">
                      {currentSection.charAt(0).toUpperCase() + currentSection.slice(1)} ({getCurrentData().length})
                    </h2>
                    <button
                      onClick={() => {
                        setShowForm(true);
                        setEditingItem(null);
                        setFormData({});
                      }}
                      className="bg-indigo-600 hover:bg-indigo-700 text-white font-bold py-2 px-4 rounded"
                    >
                      Agregar {getSectionTitle()}
                    </button>
                  </div>
                  
                  {loading ? (
                    <div className="flex justify-center py-8">
                      <div className="animate-spin rounded-full h-8 w-8 border-b-2 border-indigo-600"></div>
                    </div>
                  ) : getCurrentData().length === 0 ? (
                    <div className="text-center py-8 text-gray-500">
                      No hay {currentSection} registradas.
                    </div>
                  ) : (
                    <div className="overflow-x-auto">
                      {renderTable()}
                    </div>
                  )}
                </div>
              </div>
            )}
          </div>
        </div>
      </div>

      {/* Modal de confirmación para eliminación */}
      <ConfirmDialog
        isOpen={showConfirmDialog}
        onClose={handleCancelDelete}
        onConfirm={handleConfirmDelete}
        title="Confirmar Eliminación"
        message={`¿Está seguro de que desea eliminar "${itemToDelete?.nombre || itemToDelete?.ciudad || itemToDelete?.nombre_institucion || 'este elemento'}"? Esta acción no se puede deshacer.`}
        confirmText="Eliminar"
        cancelText="Cancelar"
      />
    </div>
  );
};

export default SystemConfig;
