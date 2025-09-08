import { useState, useEffect } from 'react';
import axios from 'axios';
import ConfirmDialog from './ConfirmDialog';

const EstudiantesManager = ({ onBack }) => {
  // URL base de la API
  const API_BASE_URL = 'http://localhost:3000';
  
  const [estudiantes, setEstudiantes] = useState([]);
  const [ciudades, setCiudades] = useState([]);
  const [instituciones, setInstituciones] = useState([]);
  const [provincias, setProvincias] = useState([]);
  const [loading, setLoading] = useState(true);
  const [showForm, setShowForm] = useState(false);
  const [editingEstudiante, setEditingEstudiante] = useState(null);
  const [error, setError] = useState('');
  const [success, setSuccess] = useState('');
  
  // Estados para el modal de confirmación
  const [showConfirmDialog, setShowConfirmDialog] = useState(false);
  const [estudianteToDelete, setEstudianteToDelete] = useState(null);

  const [formData, setFormData] = useState({
    // Datos de la persona
    nombre: '',
    fecha_nacimiento: '',
    correo: '',
    telefono: '',
    cedula: '',
    // Datos del estudiante
    institucion_id: '',
    ciudad_id: '',
    especialidad: ''
  });

  // Cargar datos iniciales
  useEffect(() => {
    loadInitialData();
  }, []);

  const loadInitialData = async () => {
    try {
      const [estudiantesRes, ciudadesRes, institucionesRes, provinciasRes] = await Promise.all([
        axios.get(`${API_BASE_URL}/estudiantes`),
        axios.get(`${API_BASE_URL}/ciudades`),
        axios.get(`${API_BASE_URL}/instituciones`),
        axios.get(`${API_BASE_URL}/provincias`)
      ]);

      setEstudiantes(estudiantesRes.data);
      setCiudades(ciudadesRes.data);
      setInstituciones(institucionesRes.data);
      setProvincias(provinciasRes.data);
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
      let personaId;

      if (editingEstudiante) {
        // Si estamos editando, actualizar la persona existente
        const personaData = {
          nombre: formData.nombre,
          fecha_nacimiento: formData.fecha_nacimiento ? formData.fecha_nacimiento + 'T00:00:00Z' : null,
          correo: formData.correo || null,
          telefono: formData.telefono || null,
          cedula: formData.cedula
        };

        console.log('Actualizando datos de persona:', personaData); // Para debug
        console.log('Estudiante siendo editado:', editingEstudiante); // Para debug
        
        // Obtener el persona_id del estudiante
        const personaIdToUpdate = editingEstudiante.persona_id || editingEstudiante.persona?.ID;
        console.log('Persona ID a actualizar:', personaIdToUpdate); // Para debug
        
        if (!personaIdToUpdate) {
          throw new Error('No se pudo obtener el ID de la persona para actualizar');
        }
        
        // Actualizar la persona existente
        await axios.put(`${API_BASE_URL}/personas/${personaIdToUpdate}`, personaData);
        personaId = personaIdToUpdate;
      } else {
        // Si estamos creando, crear nueva persona
        const personaData = {
          nombre: formData.nombre,
          fecha_nacimiento: formData.fecha_nacimiento ? formData.fecha_nacimiento + 'T00:00:00Z' : null,
          correo: formData.correo || null,
          telefono: formData.telefono || null,
          cedula: formData.cedula
        };

        console.log('Enviando datos de persona:', personaData); // Para debug

        const personaResponse = await axios.post(`${API_BASE_URL}/personas`, personaData);
        personaId = personaResponse.data.ID;
      }

      // Crear o actualizar el estudiante
      const estudianteData = {
        persona_id: personaId,
        institucion_id: parseInt(formData.institucion_id),
        ciudad_id: parseInt(formData.ciudad_id),
        especialidad: formData.especialidad
      };

      if (editingEstudiante) {
        await axios.put(`${API_BASE_URL}/estudiantes/${editingEstudiante.ID}`, estudianteData);
        setSuccess('Estudiante actualizado exitosamente');
      } else {
        await axios.post(`${API_BASE_URL}/estudiantes`, estudianteData);
        setSuccess('Estudiante registrado exitosamente');
      }

      // Recargar la lista
      await loadInitialData();
      
      // Resetear formulario
      setFormData({
        nombre: '',
        fecha_nacimiento: '',
        correo: '',
        telefono: '',
        cedula: '',
        institucion_id: '',
        ciudad_id: '',
        especialidad: ''
      });
      setShowForm(false);
      setEditingEstudiante(null);

    } catch (err) {
      console.error('Error completo:', err); // Para debug
      console.error('Response data:', err.response?.data); // Para debug
      console.error('Response status:', err.response?.status); // Para debug
      console.error('Request config:', err.config); // Para debug
      
      let errorMessage = 'Error al guardar el estudiante: ';
      if (err.response?.data?.error) {
        errorMessage += err.response.data.error;
      } else if (err.response?.status === 400) {
        errorMessage += 'Datos inválidos. Verifique que todos los campos requeridos estén llenos correctamente.';
      } else if (err.response?.status === 500) {
        errorMessage += 'Error interno del servidor. Verifique que no haya datos duplicados (cédula/correo).';
      } else {
        errorMessage += err.message;
      }
      setError(errorMessage);
    } finally {
      setLoading(false);
    }
  };

  const handleEdit = (estudiante) => {
    console.log('Editando estudiante:', estudiante); // Para debug
    setEditingEstudiante(estudiante);
    setFormData({
      nombre: estudiante.persona?.nombre || '',
      fecha_nacimiento: estudiante.persona?.fecha_nacimiento ? estudiante.persona.fecha_nacimiento.split('T')[0] : '',
      correo: estudiante.persona?.correo || '',
      telefono: estudiante.persona?.telefono || '',
      cedula: estudiante.persona?.cedula || '',
      institucion_id: estudiante.institucion_id?.toString() || '',
      ciudad_id: estudiante.ciudad_id?.toString() || '',
      especialidad: estudiante.especialidad || ''
    });
    setShowForm(true);
  };

  const handleDeleteClick = (estudiante) => {
    setEstudianteToDelete(estudiante);
    setShowConfirmDialog(true);
  };

  const handleConfirmDelete = async () => {
    if (estudianteToDelete) {
      try {
        await axios.delete(`${API_BASE_URL}/estudiantes/${estudianteToDelete.ID}`);
        setSuccess('Estudiante eliminado exitosamente');
        await loadInitialData();
      } catch (err) {
        setError('Error al eliminar el estudiante: ' + (err.response?.data?.error || err.message));
      }
    }
    setShowConfirmDialog(false);
    setEstudianteToDelete(null);
  };

  const handleCancelDelete = () => {
    setShowConfirmDialog(false);
    setEstudianteToDelete(null);
  };

  const getCiudadNombre = (ciudadId) => {
    const ciudad = ciudades.find(c => c.ID === ciudadId);
    return ciudad ? ciudad.ciudad : 'N/A';
  };

  const getInstitucionNombre = (institucionId) => {
    const institucion = instituciones.find(i => i.ID === institucionId);
    return institucion ? institucion.nombre : 'N/A';
  };

  if (loading && !showForm) {
    return (
      <div className="min-h-screen flex items-center justify-center">
        <div className="animate-spin rounded-full h-32 w-32 border-b-2 border-indigo-600"></div>
      </div>
    );
  }

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
                Gestión de Estudiantes
              </h1>
            </div>
            <button
              onClick={() => {
                setShowForm(!showForm);
                setEditingEstudiante(null);
                setFormData({
                  nombre: '',
                  fecha_nacimiento: '',
                  correo: '',
                  telefono: '',
                  cedula: '',
                  institucion_id: '',
                  ciudad_id: '',
                  especialidad: ''
                });
              }}
              className="bg-indigo-600 hover:bg-indigo-700 text-white font-bold py-2 px-4 rounded"
            >
              {showForm ? 'Cancelar' : 'Nuevo Estudiante'}
            </button>
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

      {/* Formulario */}
      {showForm && (
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 pt-6">
          <div className="bg-white shadow rounded-lg p-6">
            <h2 className="text-xl font-bold mb-4">
              {editingEstudiante ? 'Editar Estudiante' : 'Registrar Nuevo Estudiante'}
            </h2>
            <form onSubmit={handleSubmit} className="grid grid-cols-1 md:grid-cols-2 gap-4">
              {/* Datos de la persona */}
              <div className="md:col-span-2">
                <h3 className="text-lg font-semibold mb-2">Datos Personales</h3>
              </div>
              
              <div>
                <label className="block text-sm font-medium text-gray-700">Nombre Completo</label>
                <input
                  type="text"
                  name="nombre"
                  value={formData.nombre}
                  onChange={handleInputChange}
                  required
                  className="mt-1 block w-full border border-gray-300 rounded-md shadow-sm py-2 px-3 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500"
                />
              </div>

              <div>
                <label className="block text-sm font-medium text-gray-700">Cédula</label>
                <input
                  type="text"
                  name="cedula"
                  value={formData.cedula}
                  onChange={handleInputChange}
                  required
                  className="mt-1 block w-full border border-gray-300 rounded-md shadow-sm py-2 px-3 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500"
                />
              </div>

              <div>
                <label className="block text-sm font-medium text-gray-700">Fecha de Nacimiento</label>
                <input
                  type="date"
                  name="fecha_nacimiento"
                  value={formData.fecha_nacimiento}
                  onChange={handleInputChange}
                  className="mt-1 block w-full border border-gray-300 rounded-md shadow-sm py-2 px-3 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500"
                />
              </div>

              <div>
                <label className="block text-sm font-medium text-gray-700">Correo Electrónico</label>
                <input
                  type="email"
                  name="correo"
                  value={formData.correo}
                  onChange={handleInputChange}
                  className="mt-1 block w-full border border-gray-300 rounded-md shadow-sm py-2 px-3 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500"
                />
              </div>

              <div>
                <label className="block text-sm font-medium text-gray-700">Teléfono</label>
                <input
                  type="tel"
                  name="telefono"
                  value={formData.telefono}
                  onChange={handleInputChange}
                  className="mt-1 block w-full border border-gray-300 rounded-md shadow-sm py-2 px-3 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500"
                />
              </div>

              {/* Datos del estudiante */}
              <div className="md:col-span-2 mt-4">
                <h3 className="text-lg font-semibold mb-2">Datos Académicos</h3>
              </div>

              <div>
                <label className="block text-sm font-medium text-gray-700">Institución</label>
                <select
                  name="institucion_id"
                  value={formData.institucion_id}
                  onChange={handleInputChange}
                  required
                  className="mt-1 block w-full border border-gray-300 rounded-md shadow-sm py-2 px-3 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500"
                >
                  <option value="">Seleccione una institución</option>
                  {instituciones.map((institucion) => (
                    <option key={institucion.ID} value={institucion.ID}>
                      {institucion.nombre}
                    </option>
                  ))}
                </select>
              </div>

              <div>
                <label className="block text-sm font-medium text-gray-700">Ciudad</label>
                <select
                  name="ciudad_id"
                  value={formData.ciudad_id}
                  onChange={handleInputChange}
                  required
                  className="mt-1 block w-full border border-gray-300 rounded-md shadow-sm py-2 px-3 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500"
                >
                  <option value="">Seleccione una ciudad</option>
                  {ciudades.map((ciudad) => (
                    <option key={ciudad.ID} value={ciudad.ID}>
                      {ciudad.ciudad}
                    </option>
                  ))}
                </select>
              </div>

              <div>
                <label className="block text-sm font-medium text-gray-700">Especialidad</label>
                <input
                  type="text"
                  name="especialidad"
                  value={formData.especialidad}
                  onChange={handleInputChange}
                  className="mt-1 block w-full border border-gray-300 rounded-md shadow-sm py-2 px-3 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500"
                />
              </div>

              <div className="md:col-span-2">
                <button
                  type="submit"
                  disabled={loading}
                  className="w-full bg-indigo-600 hover:bg-indigo-700 text-white font-bold py-2 px-4 rounded disabled:opacity-50"
                >
                  {loading ? 'Guardando...' : (editingEstudiante ? 'Actualizar Estudiante' : 'Registrar Estudiante')}
                </button>
              </div>
            </form>
          </div>
        </div>
      )}

      {/* Lista de estudiantes */}
      {!showForm && (
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 pt-6">
          <div className="bg-white shadow rounded-lg overflow-hidden">
            <div className="px-4 py-5 sm:p-6">
              <h3 className="text-lg leading-6 font-medium text-gray-900 mb-4">
                Lista de Estudiantes ({estudiantes.length})
              </h3>
              
              {estudiantes.length === 0 ? (
                <p className="text-gray-500">No hay estudiantes registrados.</p>
              ) : (
                <div className="overflow-x-auto">
                  <table className="min-w-full divide-y divide-gray-200">
                    <thead className="bg-gray-50">
                      <tr>
                        <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                          Nombre
                        </th>
                        <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                          Cédula
                        </th>
                        <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                          Institución
                        </th>
                        <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                          Ciudad
                        </th>
                        <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                          Especialidad
                        </th>
                        <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                          Acciones
                        </th>
                      </tr>
                    </thead>
                    <tbody className="bg-white divide-y divide-gray-200">
                      {estudiantes.map((estudiante) => (
                        <tr key={estudiante.ID}>
                          <td className="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
                            {estudiante.persona?.nombre || 'N/A'}
                          </td>
                          <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                            {estudiante.persona?.cedula || 'N/A'}
                          </td>
                          <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                            {getInstitucionNombre(estudiante.institucion_id)}
                          </td>
                          <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                            {getCiudadNombre(estudiante.ciudad_id)}
                          </td>
                          <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                            {estudiante.especialidad || 'N/A'}
                          </td>
                          <td className="px-6 py-4 whitespace-nowrap text-sm font-medium space-x-2">
                            <button
                              onClick={() => handleEdit(estudiante)}
                              className="inline-flex items-center px-3 py-1.5 border border-transparent text-xs font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 transition-colors duration-200"
                            >
                              <svg className="w-4 h-4 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z" />
                              </svg>
                              Editar
                            </button>
                            <button
                              onClick={() => handleDeleteClick(estudiante)}
                              className="inline-flex items-center px-3 py-1.5 border border-transparent text-xs font-medium rounded-md text-white bg-red-600 hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-red-500 transition-colors duration-200"
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
                </div>
              )}
            </div>
          </div>
        </div>
      )}

      {/* Modal de confirmación para eliminación */}
      <ConfirmDialog
        isOpen={showConfirmDialog}
        onClose={handleCancelDelete}
        onConfirm={handleConfirmDelete}
        title="Confirmar Eliminación"
        message={`¿Está seguro de que desea eliminar al estudiante "${estudianteToDelete?.persona?.nombre || 'este estudiante'}"? Esta acción no se puede deshacer.`}
        confirmText="Eliminar"
        cancelText="Cancelar"
      />
    </div>
  );
};

export default EstudiantesManager;
